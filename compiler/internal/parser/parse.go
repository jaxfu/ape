/*
currently redesigning;
using steps that each do context free processing
for their specific step type (i.e. key, value)
*/
package parser

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/golp/list"
)

func Parse(tokens []shared.Token, prealloc uint) (
	shared.Ast,
	[]error,
	error,
) {
	toks := list.Wrap(tokens)
	ast := make(shared.Ast, 0, prealloc)
	errors := make([]error, 0, prealloc/4)
	indentType := INDENT_UNKOWN
	counter := NodeCounter{}

	for {
		var step, next Step = stepEntry, stepEntry
		var err error
		nb := NewNodeBuilder(indentType)

		// parse node
		for {
			nb, next, err = step.process(&toks, nb)
			if err != nil || next == nil {
				break
			}
			step = next
			indentType = nb.IndentType
		}
		// process, append node
		node, err := nb.process(&counter)
		if err != nil {
			fmt.Printf("error casting node: %+v\n", nb)
		}
		ast = append(ast, node)

		// seek terminator or comment
		res, tok := seekCommentOrTerminator(&toks)
		switch res {
		case RESULT_SUCCESS:
			// if success, curr is comment, so move back
			// to be moved forward again at end of loop
			tok = toks.Move(-1)
		case RESULT_UNEXPECTED:
			err = shared.NewSyntaxError(tok.Position, "unexpected token")
		}

		// TODO: handle err
		if err != nil {
		}

		// advance if not EOF
		if toks.Look(1).Type == shared.TOKEN_EOF {
			break
		} else {
			toks.Move(1)
		}
	}

	return ast, errors, nil
}

// seeks until anything other than space
func seekNextElem(toks *list.List[shared.Token]) (
	Result,
	shared.Token,
) {
	tok := toks.Curr()
	for ; ; tok = toks.Move(1) {
		if isIn(tok.Type, spacerMap) {
			continue
		} else if tok.Type == shared.TOKEN_EOF {
			return RESULT_EOF, shared.Token{}
		} else if isIn(tok.Type, nodeTerminatorMap) {
			return RESULT_TERMINATOR, shared.Token{}
		} else {
			break
		}
	}

	return RESULT_SUCCESS, tok
}

func seekTokType(
	toks *list.List[shared.Token],
	tt shared.TokenType,
) (Result, shared.Token) {
	res, tok := seekNextElem(toks)
	if res == RESULT_SUCCESS {
		if tok.Type == tt {
			return RESULT_SUCCESS, tok
		} else {
			return RESULT_UNEXPECTED, tok
		}
	}

	return res, tok
}

func seekMapMember(
	toks *list.List[shared.Token],
	tm tokenTypesMap,
) (Result, shared.Token) {
	res, tok := seekNextElem(toks)
	if res == RESULT_SUCCESS {
		if isIn(tok.Type, tm) {
			return RESULT_SUCCESS, tok
		} else {
			return RESULT_UNEXPECTED, tok
		}
	}

	return res, tok
}

func seekCommentOrTerminator(toks *list.List[shared.Token]) (
	Result,
	shared.Token,
) {
	res, tok := seekTokType(toks, shared.TOKEN_COMMENT_SYM)
	if res == RESULT_SUCCESS {
		return RESULT_SUCCESS, tok
	}

	return res, tok
}

func seekValueEntry(toks *list.List[shared.Token]) (
	Result,
	shared.Token,
) {
	return seekMapMember(toks, valueMap)
}

func parseSymbol(
	toks *list.List[shared.Token],
) string {
	tok := toks.Curr()
	symbol := tok.Content

	for tok = toks.Move(1); ; tok = toks.Move(1) {
		if tok.Type != shared.TOKEN_SYMBOL {
			break
		}
		symbol = symbol + tok.Content
	}

	return symbol
}

func isConstraint(tok shared.Token) bool {
	return (tok.Type == shared.TOKEN_SYMBOL &&
		tok.Content == shared.SYMBOL_DECLARE_CONSTRAINT)
}
