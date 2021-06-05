package storage

import (
	"fmt"
)

type Storage interface {
	Put(k string, v interface{}) error
	Get(k string, v interface{}) error
	Delete(k string) error
}

func TestStorage(s Storage) error {

	type Some struct {
		A string
		B bool
		I int
		F float64
	}

	key := "foo"

	putReq := Some{
		A: "abc123",
		B: true,
		I: 1234567890,
		F: 12345.6789,
	}

	err := s.Put(key, putReq)
	if err != nil {
		return fmt.Errorf("unexpected error happened during Put: %v", err)
	}

	var getRes Some
	err = s.Get(key, &getRes)
	if err != nil {
		return fmt.Errorf("unexpected error happened during Get: %v", err)
	}

	if putReq != getRes {
		return fmt.Errorf("the value got did not equal the value put")
	}

	err = s.Delete(key)
	if err != nil {
		return fmt.Errorf("unexpected error happened during Delete: %v", err)
	}

	err = s.Get(key, nil)
	if _, ok := err.(*NotFound); !ok {
		return fmt.Errorf("Get did not to return not found error after deletion")
	}
	err = s.Delete(key)
	if _, ok := err.(*NotFound); !ok {
		return fmt.Errorf("Delete did not to return not found error after deletion")
	}
	return nil
}
