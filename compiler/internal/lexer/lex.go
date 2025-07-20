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
		// get rune from buf
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
				l.handleIndividualRune(rune)
			}

		case shared.TOKEN_IDENT:
			if identIntraRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
			} else if symbolRunes.contains(rune) {
				l.addIdent()
				l.addSymbol(rune)

			} else {
				l.handleIndividualRune(rune)
			}

		case shared.TOKEN_NUMBER:
			if numberIntraRunes.contains(rune) {
				l.StringBuilder.WriteRune(rune)
			} else {
				l.handleIndividualRune(rune)
			}

		case shared.TOKEN_STRING:
			switch rune {
			case shared.RUNE_SPACE:
				l.addString()
				l.addSpace()
			case shared.RUNE_TAB:
				l.addString()
				l.addTab()
			case shared.RUNE_CARRIAGE_RETURN:
				l.handleCR()
			case shared.RUNE_LINEFEED:
				l.addString()
				l.addNewline()
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
	l.StringBuilder.Reset()

	l.appendToken(
		shared.NewToken(
			shared.TOKEN_IDENT,
			cnt,
			shared.Position{
				Line: l.Position.Line,
				Col:  l.Position.Col - uint(len(cnt)),
			},
		),
	)
	l.Status = shared.TOKEN_UNDEFINED
}

func (l *Lexer) addNumber() {
	cnt := l.StringBuilder.String()
	l.StringBuilder.Reset()

	l.appendToken(
		shared.NewToken(
			shared.TOKEN_NUMBER,
			cnt,
			shared.Position{
				Line: l.Position.Line,
				Col:  l.Position.Col - uint(len(cnt)),
			},
		),
	)
	l.Status = shared.TOKEN_UNDEFINED
}

func (l *Lexer) addString() {
	cnt := l.StringBuilder.String()
	l.StringBuilder.Reset()

	l.appendToken(
		shared.NewToken(
			shared.TOKEN_STRING,
			cnt,
			shared.Position{
				Line: l.Position.Line,
				Col:  l.Position.Col - uint(len(cnt)),
			},
		),
	)
	l.Status = shared.TOKEN_UNDEFINED
}

func (l *Lexer) addSymbol(r rune) {
	l.createAndAppendToken(shared.TOKEN_SYMBOL, string(r))
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

func (l *Lexer) addFromStatus() {
	switch l.Status {
	case shared.TOKEN_IDENT:
		l.addIdent()
	case shared.TOKEN_STRING:
		l.addString()
	case shared.TOKEN_NUMBER:
		l.addNumber()
	}
}

func (l *Lexer) handleIndividualRune(r rune) {
	switch r {
	case shared.RUNE_SPACE:
		l.addFromStatus()
		l.addSpace()
		l.Status = shared.TOKEN_UNDEFINED
	case shared.RUNE_TAB:
		l.addFromStatus()
		l.addTab()
		l.Status = shared.TOKEN_UNDEFINED
	case shared.RUNE_CARRIAGE_RETURN:
		l.handleCR()
	case shared.RUNE_LINEFEED:
		l.addFromStatus()
		l.addNewline()
		l.Status = shared.TOKEN_UNDEFINED
	default: // unmatched, change status to string
		l.StringBuilder.WriteRune(r)
		l.Status = shared.TOKEN_STRING
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
