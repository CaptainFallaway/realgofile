package filestorage

import "io"

type FileStorage interface {
	NewTransaction() (io.WriteCloser, error)
}
