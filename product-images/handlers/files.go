package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/mwazovzky/microservices-introduction/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	logger hclog.Logger
	store  files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, logger: l}
}

// curl -vv localhost:9090/images/1/test.png -X POST --data-binary @gopher.png
// ServeHTTP implements the http.Handler interface
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filename := vars["filename"]

	f.logger.Info("Handle POST", "id", id, "filename", filename)

	// check
	if id == "" || filename == "" {
		f.invalidURI(rw, r.URL.String())
	}

	f.saveFile(id, filename, rw, r)
}

func (f *Files) invalidURI(rw http.ResponseWriter, uri string) {
	f.logger.Error("Invalid path", "path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.logger.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.logger.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
