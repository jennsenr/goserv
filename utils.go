package goserv

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) getFullPath(httpMethod, path string) string {
	var fullPath string
	if s.basePath != "" {
		fullPath = s.basePath + path
	} else {
		fullPath = path
	}

	return fmt.Sprintf("%s %s", httpMethod, strings.TrimSuffix(fullPath, "/"))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

		next.ServeHTTP(w, r)
	})
}
