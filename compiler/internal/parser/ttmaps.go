package parser

import (
	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/golp/list"
)

// entry
var ttmNodeEntry = newTokenTypesMap(
	shared.TOKEN_SYMBOL,
	shared.TOKEN_IDENT,
)

var ttmCommentEntry = newTokenTypesMap(
	shared.TOKEN_COMMENT_SYM,
)

var ttmIndent = newTokenTypesMap(
	shared.TOKEN_SPACE,
	shared.TOKEN_TAB,
)

// key
var ttmKeyEntry = newTokenTypesMap(
	shared.TOKEN_IDENT,
	shared.TOKEN_SYMBOL,
)

// value
var ttmValue = newTokenTypesMap(
	shared.TOKEN_IDENT,
	shared.TOKEN_SYMBOL,
	shared.TOKEN_STRING,
	shared.TOKEN_NUMBER,
)

// nodes
var ttmNodeTerminator = newTokenTypesMap(
	shared.TOKEN_NEWLINE,
	shared.TOKEN_EOF,
)

var ttmAfterCompleteNode = newTokenTypesMap(
	shared.TOKEN_NEWLINE,
	shared.TOKEN_COMMENT_SYM,
)

var ttmSpacer = newTokenTypesMap(
	shared.TOKEN_SPACE,
	shared.TOKEN_TAB,
)

type tokenTypesMap map[shared.TokenType]struct{}

func newTokenTypesMap(toks ...shared.TokenType) tokenTypesMap {
	m := map[shared.TokenType]struct{}{}
	for _, t := range toks {
		m[t] = struct{}{}
	}
	return m
}

func (tm tokenTypesMap) contains(toktype shared.TokenType) bool {
	_, ok := tm[toktype]
	return ok
}

func (tm tokenTypesMap) seekMember(
	toks *list.List[shared.Token],
) (Result, shared.Token) {
	res, tok := seekNextElem(toks)
	if res == RESULT_SUCCESS {
		if tm.contains(tok.Type) {
			return RESULT_SUCCESS, tok
		} else {
			return RESULT_UNEXPECTED, tok
		}
	}

	return res, tok
}
