package database

import (
	"context"
	"sync"
	"time"

	log "github.com/siruspen/logrus"

	"github.com/secmohammed/restaurant-management/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
)

type Database interface {
	OpenCollection(name string) *mongo.Collection
}

type database struct {
	client *mongo.Client
	dbName string
}

func NewDatabaseConnection(config *config.Config) Database {
	once.Do(func() {
		db := config.Get().GetString("db.url")
		connection, err := mongo.NewClient(options.Client().ApplyURI(db))
		client = connection
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Connected to MongoDB")
	})
	return &database{client: client, dbName: config.Get().GetString("db.name")}
}

func (d *database) OpenCollection(collectionName string) *mongo.Collection {
	collection := d.client.Database(d.dbName).Collection(collectionName)
	return collection
}
