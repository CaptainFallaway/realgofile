package filestorage

import "io"

type FileStorage interface {
	Delete(uid string) error
	NewWriteTransaction(uid string) (io.WriteCloser, error)
	NewReadTransaction(uid string) (io.ReadCloser, error)
}
