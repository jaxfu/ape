package lexer

import (
	"fmt"
	"io"

	"github.com/jaxfu/ape/compiler/internal/shared"
)

type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

func (l *Lexer) Lex(rdr RuneReader, prealloc uint) ([]shared.Token, error) {
	l.setup(prealloc)

	for ; ; l.Position.Col++ {
		rune, _, err := rdr.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading rune: %+v", err)
		}

		// top-level switch based on current token being built
		switch l.Status {

		case shared.TOKEN_UNDEFINED:
			// check groups first
			// can allow utf-8 with unicode.IsLetter()
			if identEntryRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
				l.Status = shared.TOKEN_IDENT
			} else if numberEntryRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
				l.Status = shared.TOKEN_NUMBER
			} else if symbolRunes.contains(rune) {
				l.addSymbol(rune)
			} else { // check against individual runes
				switch rune {
				case shared.RUNE_SPACE:
					l.addSpace()
				case shared.RUNE_TAB:
					l.addTab()
				case shared.RUNE_SYM_COMMENT:
					l.addComment()
				case shared.RUNE_CARRIAGE_RETURN:
					l.handleCR()
				case shared.RUNE_LINEFEED:
					l.addNewline()
				default: // unmatched, start string literal
					// TODO: unmatched rune sanitization
					l.Status = shared.TOKEN_STRING
					l.StringBuilder.WriteRune(rune)
				}
			}

		case shared.TOKEN_IDENT:
			if identIntraRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
			} else if symbolRunes.contains(rune) {
				l.addIdent()
				l.addSymbol(rune)
				l.Status = shared.TOKEN_UNDEFINED
			} else {
				switch rune {
				case shared.RUNE_SPACE:
					l.addIdent()
					l.addSpace()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_TAB:
					l.addIdent()
					l.addTab()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_SYM_COMMENT:
					l.addIdent()
					l.addComment()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_CARRIAGE_RETURN:
					l.handleCR()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_LINEFEED:
					l.addIdent()
					l.addNewline()
					l.Status = shared.TOKEN_UNDEFINED
				default: // unmatched, change status to string
					l.Status = shared.TOKEN_STRING
					l.StringBuilder.WriteRune(rune)
				}
			}

		case shared.TOKEN_NUMBER:
			if numberIntraRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
			} else {
				switch rune {
				case shared.RUNE_SPACE:
					l.addNumber()
					l.addSpace()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_TAB:
					l.addNumber()
					l.addTab()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_POUND:
					l.addNumber()
					l.addComment()
					l.Status = shared.TOKEN_UNDEFINED
				case shared.RUNE_CARRIAGE_RETURN:
					l.handleCR()
				case shared.RUNE_LINEFEED:
					l.addNumber()
					l.addNewline()
					l.Status = shared.TOKEN_UNDEFINED
				default: // unmatched, change status to string
					l.Status = shared.TOKEN_STRING
					l.StringBuilder.WriteRune(rune)
				}
			}

		case shared.TOKEN_STRING:
			switch rune {
			case shared.RUNE_SPACE:
				l.addString()
				l.addSpace()
				l.Status = shared.TOKEN_UNDEFINED
			case shared.RUNE_TAB:
				l.addString()
				l.addTab()
				l.Status = shared.TOKEN_UNDEFINED
			case shared.RUNE_POUND:
				l.addString()
				l.addComment()
				l.Status = shared.TOKEN_UNDEFINED
			case shared.RUNE_CARRIAGE_RETURN:
				l.handleCR()
			case shared.RUNE_LINEFEED:
				l.addString()
				l.addNewline()
				l.Status = shared.TOKEN_UNDEFINED
			default: // unmatched, write
				l.StringBuilder.WriteRune(rune)
			}
		}
	}

	// cleanup
	switch l.Status {
	case shared.TOKEN_IDENT:
		l.addIdent()
	case shared.TOKEN_STRING:
		l.addString()
	case shared.TOKEN_NUMBER:
		l.addNumber()
	}

	// add EOF token
	l.createAndAppendToken(shared.TOKEN_EOF, shared.TOKEN_EOF)

	return l.Tokens, nil
}

func (l *Lexer) addIdent() {
	cnt := l.StringBuilder.String()
	tok := shared.NewToken(
		shared.TOKEN_IDENT,
		cnt,
		shared.Position{
			Line: l.Position.Line,
			Col:  l.Position.Col - uint(len(cnt)),
		},
	)
	l.appendToken(tok)
	l.StringBuilder.Reset()
	l.Status = shared.TOKEN_UNDEFINED
}

func (l *Lexer) addNumber() {
	cnt := l.StringBuilder.String()
	tok := shared.NewToken(
		shared.TOKEN_NUMBER,
		cnt,
		shared.Position{
			Line: l.Position.Line,
			Col:  l.Position.Col - uint(len(cnt)),
		},
	)
	l.appendToken(tok)
	l.StringBuilder.Reset()
}

func (l *Lexer) addString() {
	cnt := l.StringBuilder.String()
	tok := shared.NewToken(
		shared.TOKEN_STRING,
		cnt,
		shared.Position{
			Line: l.Position.Line,
			Col:  l.Position.Col - uint(len(cnt)),
		},
	)
	l.appendToken(tok)
	l.StringBuilder.Reset()
}

func (l *Lexer) addSymbol(rune rune) {
	l.createAndAppendToken(shared.TOKEN_SYMBOL, string(rune))
}

func (l *Lexer) addSpace() {
	l.createAndAppendToken(shared.TOKEN_SPACE, shared.SYMBOL_SPACE)
}

func (l *Lexer) addTab() {
	l.createAndAppendToken(shared.TOKEN_TAB, shared.SYMBOL_TAB)
}

func (l *Lexer) handleCR() {
	if l.HasFoundLineEnd {
		if l.IsUnix {
			l.Errors = append(l.Errors, shared.NewSyntaxError(l.Position, "inconsistent line endings"))
		}
	} else {
		l.HasFoundLineEnd = true
		l.IsUnix = false
	}
}

func (l *Lexer) addNewline() {
	if !l.HasFoundLineEnd {
		l.HasFoundLineEnd = true
		l.IsUnix = true
	}
	l.createAndAppendToken(shared.TOKEN_NEWLINE, shared.SYMBOL_NEWLINE)
	l.nextline()
}

func (l *Lexer) addComment() {
	l.createAndAppendToken(shared.TOKEN_COMMENT_SYM, shared.SYMBOL_COMMENT)
}
