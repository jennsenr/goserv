package goserv

import (
	"context"
	"encoding/json"
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(request *http.Request) *Request {
	return &Request{Request: request}
}

func (r *Request) GetPathParam(key string) string {
	return r.PathValue(key)
}

func (r *Request) GetQueryParam(key string) string {
	return r.URL.Query().Get(key)
}

func (r *Request) Bind(t any) error {
	return json.NewDecoder(r.Request.Body).Decode(&t)
}

func (r *Request) Context() context.Context {
	return r.Request.Context()
}

func (r *Request) SetContextValue(key string, value any) {
	newContext := context.WithValue(r.Request.Context(), key, value)
	r.Request = r.Request.WithContext(newContext)
}

func (r *Request) GetContextValue(key string) any {
	return r.Request.Context().Value(key)
}

func (r *Request) GetHeaderValue(key string) string {
	return r.Header.Get(key)
}

func (r *Request) Path() string {
	return r.Request.URL.Path
}
