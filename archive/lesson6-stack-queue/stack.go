package main

type Stack[T any] struct {
	nodes []T
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.nodes) < 1 {
		var dummy T
		return dummy, false
	}

	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node, true
}

func (s *Stack[T]) Push(node T) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.nodes) < 1 {
		var dummy T
		return dummy, false
	}
	return s.nodes[len(s.nodes)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.nodes)
}
