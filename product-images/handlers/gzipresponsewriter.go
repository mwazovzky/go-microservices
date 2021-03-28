package handlers

import (
	"compress/gzip"
	"net/http"
)

type GzipResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewGzipResponseWriter(rw http.ResponseWriter) *GzipResponseWriter {
	gw := gzip.NewWriter(rw)
	return &GzipResponseWriter{rw, gw}
}

func (wr *GzipResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *GzipResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *GzipResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

func (wr *GzipResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
