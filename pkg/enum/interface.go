package enum

import "strings"

type Enum[T ~string, L any] struct {
	TypeList L
	MatchMap map[string]T
}

func (e Enum[T, L]) Types() L {
	return e.TypeList
}

func (e Enum[T, L]) Match(src string) T {
	src = strings.ToLower(src)
	src = strings.TrimSpace(src)

	found, ok := e.MatchMap[src]
	if !ok {
		return "UNDEFINED"
	}

	return found
}
