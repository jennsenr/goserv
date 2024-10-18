package goserv

import "errors"

var (
	ErrNotFound     = errors.New("not_found")
	ErrInternal     = errors.New("internal_error")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad_request")
)
