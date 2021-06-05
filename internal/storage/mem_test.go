package storage_test

import (
	"testing"

	"github.com/hyqe/flint/internal/storage"
)

func TestMemory(t *testing.T) {
	mem := &storage.Memory{}
	if err := storage.TestStorage(mem); err != nil {
		t.Fatal(err)
	}
}
