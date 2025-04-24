package bus

func NewChannel[T any](bufsize int) chan T {
	var newchan chan T
	if bufsize < 1 {
		newchan = make(chan T)
	} else {
		newchan = make(chan T, bufsize)
	}
	return newchan
}
