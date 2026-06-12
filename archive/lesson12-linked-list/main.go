package main

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	head *node[T]
	tail *node[T]
}

func (l *List[T]) PushHead(v T) {
	newHead := &node[T]{value: v, next: l.head}
	l.head = newHead
	if l.tail == nil {
		l.tail = newHead
	}
}

func (l *List[T]) PushTail(v T) {
	tail := l.tail
	newTail := &node[T]{value: v, next: nil}

	if tail == nil {
		l.head = newTail
		l.tail = newTail
		return
	}
	tail.next = newTail
	l.tail = newTail
}

func (l *List[T]) PopHead() (T, bool) {
	if l.head == nil {
		var dummy T
		return dummy, false
	}

	out := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return out, true
}

func (l *List[T]) PeekHead() (T, bool) {
	if l.head == nil {
		var dummy T
		return dummy, false
	}

	return l.head.value, true
}

func (l *List[T]) Len() int {
	var length int
	for n := l.head; n != nil; n = n.next {
		length++
	}
	return length
}

func (l *List[T]) ToSlice() []T {
	var xs []T
	for n := l.head; n != nil; n = n.next {
		xs = append(xs, n.value)
	}
	return xs
}

func (l *List[T]) Iter(fn func(val T) bool) {
	for n := l.head; n != nil; n = n.next {
		if ok := fn(n.value); !ok {
			return
		}
	}
}

func (l *List[T]) Reverse() {
	if l.head == nil {
		return
	}
	l.tail = l.head
	var prev *node[T]
	curr := l.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func main() {
	l := &List[int]{}

	l.PushTail(1)
	l.PushTail(2)
	l.PushTail(3)
	l.PushHead(0)
	l.PushHead(4)
	l.PushHead(5)
	l.PushHead(6)

	println("Initial list:")
	println("  Len:", l.Len())
	for _, v := range l.ToSlice() {
		print(v, " ")
	}
	println()

	l.Iter(func(val int) bool {
		print("iter:", val, " ")
		return true
	})
	println()

	v, ok := l.PopHead()
	println("PopHead:", v, ok)
	println("After pop:", l.Len(), "items")

	l.Reverse()
	println("After reverse:")
	for _, v := range l.ToSlice() {
		print(v, " ")
	}
	println()
}
