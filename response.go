package goserv

import (
	"net/http"
)

type Response struct {
	StatusCode int     `json:"-"`
	Data       any     `json:"data"`
	Error      *string `json:"error"`
}

var emptyData = map[string]any{}

func NewDataResponse(data any) Response {
	return Response{
		StatusCode: http.StatusOK,
		Data:       data,
		Error:      nil,
	}
}

func NewErrorResponse(statusCode int, err error) Response {
	errCode := err.Error()

	return Response{
		StatusCode: statusCode,
		Data:       emptyData,
		Error:      &errCode,
	}
}

func NewEmptyResponse() Response {
	return NewDataResponse(emptyData)
}

func NewNotFoundResponse() Response {
	return NewErrorResponse(http.StatusNotFound, ErrNotFound)
}

func NewInternalErrorResponse() Response {
	return NewErrorResponse(http.StatusInternalServerError, ErrInternal)
}

func NewUnauthorizedResponse() Response {
	return NewErrorResponse(http.StatusUnauthorized, ErrUnauthorized)
}

func NewBadRequestResponse() Response {
	return NewErrorResponse(http.StatusBadRequest, ErrBadRequest)
}
