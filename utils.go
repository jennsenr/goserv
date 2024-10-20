package goserv

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) getFullPath(httpMethod, path string) string {
	var fullPath string
	
	if s.basePath != "" && s.basePath != "/" {
		fullPath = s.basePath + path
	} else {
		fullPath = path
	}

	if len(fullPath) > 1 {
		fullPath = strings.TrimSuffix(fullPath, "/")
	}

	return fmt.Sprintf("%s %s", httpMethod, fullPath)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) > 1 {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}
