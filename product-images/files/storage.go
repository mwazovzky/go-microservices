package files

import (
	"io"
	"os"
)

// Storage defines the behavior for file operations
// Implementations may be of the type local disk, or cloud storage, etc
type Storage interface {
	Index(id string) ([]string, error)
	Get(path string) (*os.File, error)
	Save(path string, file io.Reader) error
}
