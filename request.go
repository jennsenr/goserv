package goserv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (r *Request) Headers() map[string]any {
	headerBytes, err := json.Marshal(r.Request.Header)
	if err != nil {
		return make(map[string]any)
	}

	headers := make(map[string]any)

	err = json.Unmarshal(headerBytes, &headers)
	if err != nil {
		return make(map[string]any)
	}

	return headers
}

func (r *Request) BodyBytes() ([]byte, error) {
	if r.Request.Body == nil {
		return []byte{}, nil
	}

	reqBody, err := io.ReadAll(r.Request.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read request body: %w", err)
	}

	r.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	return reqBody, nil
}

func (r *Request) BodyMap() (map[string]any, error) {
	body, err := r.BodyBytes()
	if err != nil {
		return make(map[string]any), err
	}

	if len(body) == 0 {
		return make(map[string]any), nil
	}

	var bodyMap map[string]any

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return make(map[string]any), fmt.Errorf("cannot unmarshal request body: %w", err)
	}

	return bodyMap, nil
}
