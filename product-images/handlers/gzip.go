package handlers

import (
	"net/http"
	"strings"
)

type Gzip struct{}

func (g *Gzip) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			grw := NewGzipResponseWriter(rw)
			grw.Header().Set("Content-Encoding", "gzip")
			defer grw.Flush()
			next.ServeHTTP(grw, r)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
