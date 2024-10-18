package goserv

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	basePath          string
	port              string
	mux               *http.ServeMux
	globalMiddlewares []MiddlewareFunc
}

func New(basePath, port string) *Server {
	return &Server{
		basePath:          basePath,
		port:              port,
		mux:               http.NewServeMux(),
		globalMiddlewares: make([]MiddlewareFunc, 0),
	}
}

func (s *Server) Use(middlewares ...MiddlewareFunc) {
	s.globalMiddlewares = middlewares
}

func (s *Server) GET(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	s.mux.HandleFunc(s.getFullPath(http.MethodGet, path), s.applyMiddlewares(middlewares, handlerFunc))
}

func (s *Server) POST(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	s.mux.HandleFunc(s.getFullPath(http.MethodPost, path), s.applyMiddlewares(middlewares, handlerFunc))
}

func (s *Server) PUT(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	s.mux.HandleFunc(s.getFullPath(http.MethodPut, path), s.applyMiddlewares(middlewares, handlerFunc))
}

func (s *Server) PATCH(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	s.mux.HandleFunc(s.getFullPath(http.MethodPatch, path), s.applyMiddlewares(middlewares, handlerFunc))
}

func (s *Server) DELETE(path string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	s.mux.HandleFunc(s.getFullPath(http.MethodDelete, path), s.applyMiddlewares(middlewares, handlerFunc))
}

func (s *Server) Start() error {
	log.Printf("starting server on port %v\n", s.port)
	return http.ListenAndServe(":"+s.port, removeTrailingSlash(s.mux))
}

func (s *Server) applyMiddlewares(middlewares []MiddlewareFunc, handlerFunc HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{
			Request: r,
		}

		for _, middleware := range s.globalMiddlewares {
			resp := middleware(req)
			if resp != nil {
				s.writeResponse(w, *resp)
				return
			}
		}

		for _, middleware := range middlewares {
			resp := middleware(req)
			if resp != nil {
				s.writeResponse(w, *resp)
				return
			}
		}

		resp := handlerFunc(req)

		s.writeResponse(w, resp)
	}
}

func (s *Server) writeResponse(w http.ResponseWriter, resp Response) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("error writing response: %v", err)
		return
	}
}
