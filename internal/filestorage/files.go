package filestorage

import "io"

type commonStorage struct {
	basePath string
}

func NewDirectoryStorage(basePath string) FileStorage {
	return &commonStorage{basePath: basePath}
}

func (ds *commonStorage) Delete(uid string) error {
	return nil
}

func (ds *commonStorage) NewWriteTransaction(uid string) (io.WriteCloser, error) {
	return nil, nil
}

func (ds *commonStorage) NewReadTransaction(uid string) (io.ReadCloser, error) {
	return nil, nil
}
