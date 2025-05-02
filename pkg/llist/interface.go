package llist

type LinkedList[T any] struct {
	Head *Node[T]
	Curr *Node[T]
}

func (ll *LinkedList[T]) Jump(jmps int) T {
	for range jmps {
		if ll.Curr.Next != nil {
			ll.Curr = ll.Curr.Next
		}
	}

	return ll.Curr.Val
}

type Node[T any] struct {
	Prev *Node[T]
	Next *Node[T]
	Val  T
}
