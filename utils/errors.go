package utils

import (
	"errors"
	"net/http"

	"github.com/siruspen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func CreateApiError(status int, err error) (int, *ApiError) {
	logrus.Error(err.Error())
	message := err.Error()
	return status, &ApiError{
		Status:  status,
		Message: message,
	}
}

func ErrorFromDatabase(err error) (int, *ApiError) {
	switch err {
	case mongo.ErrNoDocuments:
		return CreateApiError(http.StatusNotFound, errors.New("document not found"))
	default:
		return CreateApiError(http.StatusInternalServerError, err)
	}
}
