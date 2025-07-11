package lexer

import "github.com/jaxfu/ape/compiler/internal/shared"

var symbolRunes = RuneMap{
	shared.RUNE_SYM_REFERENCE:      {},
	shared.RUNE_SYM_OPTIONAL:       {},
	shared.RUNE_SYM_START_ARRAY:    struct{}{},
	shared.RUNE_SYM_END_ARRAY:      struct{}{},
	shared.RUNE_SYM_ANON_COMPONENT: struct{}{},
}

var numberEntryRunes = RuneMap{
	// digits
	'0': {}, '1': {}, '2': {}, '3': {}, '4': {},
	'5': {}, '6': {}, '7': {}, '8': {}, '9': {},
}

var numberIntraRunes = RuneMap{
	// symbols
	'.': {},
	// digits
	'0': {}, '1': {}, '2': {}, '3': {}, '4': {},
	'5': {}, '6': {}, '7': {}, '8': {}, '9': {},
}

var identEntryRunes = RuneMap{
	'A': {}, 'B': {}, 'C': {}, 'D': {}, 'E': {}, 'F': {}, 'G': {},
	'H': {}, 'I': {}, 'J': {}, 'K': {}, 'L': {}, 'M': {}, 'N': {},
	'O': {}, 'P': {}, 'Q': {}, 'R': {}, 'S': {}, 'T': {}, 'U': {},
	'V': {}, 'W': {}, 'X': {}, 'Y': {}, 'Z': {},
	'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {},
	'h': {}, 'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {},
	'o': {}, 'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {},
	'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
}

var identIntraRunes = RuneMap{
	// symbols
	shared.RUNE_UNDERSCORE: {},
	// digits
	'0': {}, '1': {}, '2': {}, '3': {}, '4': {},
	'5': {}, '6': {}, '7': {}, '8': {}, '9': {},
	// letters
	'A': {}, 'B': {}, 'C': {}, 'D': {}, 'E': {}, 'F': {}, 'G': {},
	'H': {}, 'I': {}, 'J': {}, 'K': {}, 'L': {}, 'M': {}, 'N': {},
	'O': {}, 'P': {}, 'Q': {}, 'R': {}, 'S': {}, 'T': {}, 'U': {},
	'V': {}, 'W': {}, 'X': {}, 'Y': {}, 'Z': {},
	'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {},
	'h': {}, 'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {},
	'o': {}, 'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {},
	'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
}

var indentRunes = RuneMap{
	shared.RUNE_SPACE: {},
	shared.RUNE_TAB:   {},
}

type RuneMap map[rune]struct{}

func (m RuneMap) contains(r rune) bool {
	_, ok := m[r]
	return ok
}
