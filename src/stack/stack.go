package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the stack
type Item generic.Type

// Stack the stack of Items
type Stack struct {
	items []Item
	lock  sync.RWMutex
}

// New  creates a new stack
func (s *Stack) New() *Stack {
	s.items = []Item{}
	return s
}

// Push puts the item on the stack
func (s *Stack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

//Pop removes the item from the stack
func (s *Stack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}
