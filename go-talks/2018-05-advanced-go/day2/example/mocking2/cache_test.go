package mocking2_test

import (
	"fmt"
	"sync"
	"testing"
)

type FakeCache struct {
	mu sync.Mutex
	m  map[string][]byte

	Error error
}

func (f *FakeCache) Put(key string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.Error != nil {
		return f.Error
	}

	f.m[key] = data

	return nil
}

func (f *FakeCache) Get(key string) ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.Error != nil {
		return nil, f.Error
	}

	if v, ok := f.m[key]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("key not found")
}

func TestCache(t *testing.T) {
	tt := []struct {
		Name string
	}{}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			// ...
		})
	}
}
