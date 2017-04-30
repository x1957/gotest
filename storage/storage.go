package storage

import "fmt"
import "sync"

type Store struct {
	data map[string]bool
	lock *sync.RWMutex
}

func NewStorage() *Store {
	return &Store {
		data: make(map[string]bool),
		lock: new(sync.RWMutex),
	}
}
func (s *Store) Insert(data string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[data] = true
}

func (s *Store) GetList() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for k := range s.data {
		fmt.Println(k)
	}
}
