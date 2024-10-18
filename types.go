package goserv

type HandlerFunc func(req *Request) Response
type MiddlewareFunc func(req *Request) *Response
