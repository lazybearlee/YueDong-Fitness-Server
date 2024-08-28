package core

import (
	"net/http"
	"os"
)

type fileSystem struct {
	fs http.FileSystem
}

func (fs fileSystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}
