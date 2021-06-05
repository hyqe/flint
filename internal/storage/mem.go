package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

type Memory struct {
	db map[string]bytes.Buffer
	sync.RWMutex
}

func (m *Memory) Put(k string, v interface{}) error {
	m.init()

	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(v)
	if err != nil {
		return &Internal{Message: fmt.Sprintf("encoding error: %v", err)}
	}

	m.Lock()
	defer m.Unlock()

	m.db[k] = buff
	return nil
}

func (m *Memory) Get(k string, v interface{}) error {
	m.init()

	m.RLock()
	defer m.RUnlock()

	buff, ok := m.db[k]
	if !ok {
		return &NotFound{"not found"}
	}

	err := gob.NewDecoder(&buff).Decode(v)
	if err != nil {
		return &Internal{Message: fmt.Sprintf("decoding error: %v", err)}
	}

	return nil
}
func (m *Memory) Delete(k string) error {
	m.init()

	m.Lock()
	defer m.Unlock()
	_, ok := m.db[k]
	if ok {
		delete(m.db, k)
		return nil
	}
	return &NotFound{"not found"}
}

func (m *Memory) init() {
	m.Lock()
	defer m.Unlock()
	if m.db == nil {
		m.db = make(map[string]bytes.Buffer)
	}
}
