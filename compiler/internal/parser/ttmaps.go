package parser

import "github.com/jaxfu/ape/compiler/internal/shared"

type tokenTypesMap map[shared.TokenType]struct{}

func newTokenTypesMap(toks ...shared.TokenType) tokenTypesMap {
	m := map[shared.TokenType]struct{}{}
	for _, t := range toks {
		m[t] = struct{}{}
	}
	return m
}

// entry
var nodeEntryMap = newTokenTypesMap(
	shared.TOKEN_SYMBOL,
	shared.TOKEN_IDENT,
)

var commentEntryMap = newTokenTypesMap(
	shared.TOKEN_COMMENT_SYM,
)

var indentMap = newTokenTypesMap(
	shared.TOKEN_SPACE,
	shared.TOKEN_TAB,
)

// key
var keyEntryMap = newTokenTypesMap(
	shared.TOKEN_IDENT,
	shared.TOKEN_SYMBOL,
)

// value
var valueMap = newTokenTypesMap(
	shared.TOKEN_IDENT,
	shared.TOKEN_SYMBOL,
	shared.TOKEN_STRING,
	shared.TOKEN_NUMBER,
)

// nodes
var nodeTerminatorMap = newTokenTypesMap(
	shared.TOKEN_NEWLINE,
	shared.TOKEN_EOF,
)

var postCompleteNodeMap = newTokenTypesMap(
	shared.TOKEN_NEWLINE,
	shared.TOKEN_COMMENT_SYM,
)

var spacerMap = newTokenTypesMap(
	shared.TOKEN_SPACE,
	shared.TOKEN_TAB,
)

// returns if ttype is key in map
func isIn(toktype shared.TokenType, tokmap tokenTypesMap) bool {
	_, ok := tokmap[toktype]
	return ok
}
