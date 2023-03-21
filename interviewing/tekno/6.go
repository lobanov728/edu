package main

import "fmt"

type Store struct {
	Items int
}

func (s *Store) Incr(n int)  {
	s.Items += n
}

func (s *Store) Decr(n int)  {
	s.Items -= n
}

func (s *Store) String() string  {
	return fmt.Sprintf("%v", s.Items)
}

func main() {
	store := new(Store)
	store.Incr(10)
	store.Decr(5)
	fmt.Println(store)
}

