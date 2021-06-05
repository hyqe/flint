package storage

import "net/http"

func HttpStatus(err error) int {
	switch err.(type) {
	case *NotFound:
		return http.StatusNotFound
	case *Internal:
		return http.StatusInternalServerError
	case nil:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

type NotFound struct {
	Message string
}

func (n *NotFound) Error() string {
	return n.Message
}

type Internal struct {
	Message string
}

func (i *Internal) Error() string {
	return i.Message
}
