package storage_test

import (
	"net/http"
	"testing"

	"github.com/hyqe/flint/internal/storage"
)

func TestHttpStatus(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{err: nil}, want: http.StatusOK},
		{name: "", args: args{err: &storage.NotFound{}}, want: http.StatusNotFound},
		{name: "", args: args{err: &storage.Internal{}}, want: http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storage.HttpStatus(tt.args.err); got != tt.want {
				t.Errorf("HttpStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
