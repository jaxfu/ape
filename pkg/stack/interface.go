package stack

func NewStack[T any]() Stack[T] {
	dummy := &node[T]{
		prev: nil,
		next: nil,
		val:  *new(T),
	}

	return Stack[T]{
		head:   dummy,
		cursor: dummy,
		count:  0,
	}
}

type Stack[T any] struct {
	head   *node[T]
	cursor *node[T]
	count  int
}

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

func (s *Stack[T]) Push(t T) {
	nnode := &node[T]{
		prev: s.cursor,
		val:  t,
	}

	s.cursor.next = nnode
	s.cursor = nnode

	s.count += 1
}

func (s *Stack[T]) Pop() T {
	if s.cursor.prev == nil {
		return *new(T)
	}

	val := s.cursor.val
	s.cursor = s.cursor.prev
	s.cursor.next = nil

	s.count -= 1

	return val
}

func (s *Stack[T]) Curr() T {
	return s.cursor.val
}

func (s *Stack[T]) Size() int {
	return s.count
}
