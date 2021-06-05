package storage

import (
	"crypto/md5"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
)

type FS struct {
	Path string
	sync.RWMutex
}

func (f *FS) filePath(k string) string {
	return path.Join(f.Path, hash(k))
}

func (f *FS) Put(k string, v interface{}) error {
	key := f.filePath(k)

	f.Lock()
	defer f.Unlock()

	if hasFile(key) {
		err := deleteFile(key)
		if err != nil {
			return err
		}
	}

	return createGobFile(key, v)
}
func (f *FS) Get(k string, v interface{}) error {
	key := f.filePath(k)

	f.RLock()
	defer f.RUnlock()

	return readGobFile(key, v)
}
func (f *FS) Delete(k string) error {
	key := f.filePath(k)

	f.Lock()
	defer f.Unlock()

	if !hasFile(key) {
		return &NotFound{Message: "Not Found"}
	}

	return deleteFile(key)
}

func hash(x string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(x)))
}

func hasFile(name string) bool {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func createGobFile(name string, v interface{}) error {
	file, err := os.Create(name)
	if err != nil {
		return &Internal{Message: err.Error()}
	}
	defer file.Close()

	err = gob.NewEncoder(file).Encode(v)
	if err != nil {
		return &Internal{Message: err.Error()}
	}
	return nil
}

func deleteFile(name string) error {
	err := os.Remove(name)
	if err != nil {
		return &Internal{Message: err.Error()}
	}
	return nil
}

func readGobFile(name string, v interface{}) error {
	file, err := os.Open(name)
	if errors.Is(err, os.ErrNotExist) {
		return &NotFound{Message: err.Error()}
	}
	defer file.Close()

	err = gob.NewDecoder(file).Decode(v)
	if err != nil {
		return &Internal{Message: err.Error()}
	}
	return nil
}
