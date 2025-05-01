package lexer

import (
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
)

type Lexer struct {
	Tokens []shared.Token
	Errors []error

	Status shared.TokenType

	StringBuilder strings.Builder
	Position      shared.Position

	HasFoundLineEnd bool
	IsUnix          bool
}

type NodeBuilder struct {
	Type  shared.NodeType
	Key   string
	Value string
}

func NewLexer() Lexer {
	return Lexer{
		StringBuilder: strings.Builder{},
	}
}

// creates from type and content args, uses curr position
func (l *Lexer) createAndAppendToken(tt shared.TokenType, cnt string) {
	l.Tokens = append(
		l.Tokens,
		shared.NewToken(tt, cnt, l.Position),
	)
}

// append a premade token
func (l *Lexer) appendToken(tok shared.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func (l *Lexer) setup(prealloc uint) {
	l.StringBuilder.Reset()
	l.Position = shared.Position{
		Line: 1,
		Col:  1,
	}
	l.Tokens = make([]shared.Token, 0, prealloc)
	l.Errors = make([]error, 0, prealloc/4)
	l.Status = shared.TOKEN_UNDEFINED
}

func (l *Lexer) nextline() {
	l.StringBuilder.Reset()
	l.Position = shared.Position{
		Line: l.Position.Line + 1,
		Col:  0,
	}
	l.Status = shared.TOKEN_UNDEFINED
}

func isSymbolRune(r rune) bool {
	_, ok := symbolRunes[r]
	return ok
}

func isNumberPreRune(r rune) bool {
	_, ok := numberPreRunes[r]
	return ok
}

func isNumberPostRune(r rune) bool {
	_, ok := numberPostRunes[r]
	return ok
}

func isIdentPreRune(r rune) bool {
	_, ok := identPreRunes[r]
	return ok
}

func isIdentPostRune(r rune) bool {
	_, ok := identPostRunes[r]
	return ok
}

func isIndentRune(r rune) bool {
	_, ok := indentRunes[r]
	return ok
}
