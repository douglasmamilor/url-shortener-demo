package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ServerResponse struct {
	Err        error
	Message    string
	StatusCode int
	Context    context.Context
	Payload    interface{}
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

// Error bubbles a response error providing an implementation of the Error interface.
// It returns the error as a string.
func (r *ServerResponse) Error() string {
	return r.Err.Error()
}

// WriteJSONResponse writes data and status code to the ResponseWriter
func WriteJSONResponse(rw http.ResponseWriter, statusCode int, content []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if _, err := rw.Write(content); err != nil {
		fmt.Println("Unable to write json response")
	}
}

// RespondWithError is a utility to help with returning api errors
func RespondWithError(err error, message string, httpStatusCode int) *ServerResponse {
	var wrappedErr error
	if err != nil {
		wrappedErr = errors.Wrap(err, message)
	} else {
		wrappedErr = errors.New(message)
	}

	logrus.WithFields(logrus.Fields{
		"err": wrappedErr,
	}).Error(message)

	return &ServerResponse{
		Err:        fmt.Errorf(message),
		StatusCode: httpStatusCode,
		Message:    message,
	}
}
