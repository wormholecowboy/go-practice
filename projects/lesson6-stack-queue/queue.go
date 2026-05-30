package main

type Queue[T any] struct {
	nodes []T
}

func (q *Queue[T]) Enqueue(node T) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.nodes) < 1 {
		var dummy T
		return dummy, false
	}
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.nodes) < 1 {
		var dummy T
		return dummy, false
	}
	return q.nodes[0], true
}

func (q *Queue[T]) Size() int {
	return len(q.nodes)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.nodes) == 0
}
