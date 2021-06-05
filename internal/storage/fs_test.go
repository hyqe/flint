package storage_test

import (
	"testing"

	"github.com/hyqe/flint/internal/storage"
)

func TestFS(t *testing.T) {
	fs := &storage.FS{
		Path: ".",
	}
	if err := storage.TestStorage(fs); err != nil {
		t.Fatal(err)
	}
}
