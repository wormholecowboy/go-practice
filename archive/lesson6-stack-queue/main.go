package main

import "fmt"

func main() {

	s := Stack[int]{}
	q := Queue[int]{}

	for i := 1; i < 1000; i++ {
		s.Push(i)
		q.Enqueue(i)
	}

	for i := 1; i < 10; i++ {
		sPeek, ok := s.Peek()
		if ok {
			fmt.Println("Peek: ", sPeek)
		}
		qPeek, ok := q.Peek()
		if ok {
			fmt.Println("Peek: ", qPeek)
		}

		sPop, ok := s.Pop()
		if ok {
			fmt.Println("Pop: ", sPop)
		}
		qDequeue, ok := q.Dequeue()
		if ok {
			fmt.Println("Dequeue: ", qDequeue)
		}

	}

	fmt.Println("Is Stack Empty: ", s.IsEmpty())
	fmt.Println("Is Queue Empty: ", q.IsEmpty())
  fmt.Println("Stack size: ", s.Size())
  fmt.Println("Queue size: ", q.Size())

}
