package list

type List[T any] struct {
	cursor int
	slice  []T
}

// wrap premade slice, use as List
func Wrap[T any](slice []T) List[T] {
	return List[T]{
		slice:  slice,
		cursor: 0,
	}
}

// create List using slice of prealloc cap
func NewList[T any](prealloc uint) List[T] {
	slice := make([]T, 0, prealloc)
	return List[T]{
		slice:  slice,
		cursor: 0,
	}
}

// look without updating cursor +/-,
// overflow returns first or last
func (l *List[T]) Look(jmps int) T {
	slen := len(l.slice)

	if jmps < 0 && (l.cursor+jmps < 0) {
		return l.slice[0]
	} else if jmps > 0 && l.cursor+jmps >= slen {
		return l.slice[slen-1]
	}

	return l.slice[l.cursor+jmps]
}

// update cursor +/-, return val,
// overflow sets cursor 0 or last, returns first or last
func (l *List[T]) Move(jmps int) T {
	slen := len(l.slice)

	if jmps < 0 && (l.cursor+jmps < 0) {
		l.cursor = 0
		return l.slice[l.cursor]
	} else if jmps > 0 && l.cursor+jmps >= slen {
		l.cursor = slen - 1
		return l.slice[l.cursor]
	}

	l.cursor += jmps
	return l.slice[l.cursor]
}

// return element at cursor index in slice
func (l *List[T]) Curr() T {
	return l.slice[l.cursor]
}

// return cursor value
func (l *List[T]) Cursor() int {
	return l.cursor
}

// append to end of slice and return new cursor
func (l *List[T]) Append(el T) int {
	l.slice = append(l.slice, el)

	l.cursor++
	return l.cursor
}

func (l *List[T]) GetSlice() []T {
	return l.slice
}
