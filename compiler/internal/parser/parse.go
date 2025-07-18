/*
currently redesigning;
using steps that each do context free processing
for their specific step type (i.e. key, value)
*/
package parser

import (
	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/golp/list"
)

func Parse(tokens []shared.Token, prealloc uint) (
	[]RawNode,
	[]error,
	error,
) {
	toks := list.Wrap(tokens)
	nodes := make([]RawNode, 0, prealloc)
	// ast := make(shared.Ast, 0, prealloc)
	errors := make([]error, 0, prealloc/4)
	ctx := newParseCtx(&toks)

	for {
		var step, next Step = stepEntry, stepEntry
		var err error
		rawnode := newRawNode()

		// exit if EOF
		// if toks.Curr().Type == shared.TOKEN_EOF {
		// 	break
		// } else if toks.Curr().Type == shared.TOKEN_NEWLINE {
		// 	toks.Move(1)
		// 	continue
		// }

		// parse node
		for {
			ctx, rawnode, next, err = step.process(ctx, rawnode)
			if err != nil || next == nil {
				break
			}
			step = next
		}
		if err != nil {
			continue
		}

		// process, append node
		nodes = append(nodes, rawnode)

		// post node
		res, _ := seekNextElem(&toks)
		switch res {
		case RESULT_EOF: // move to break
		case RESULT_TERMINATOR: // if on terminator, advance
			toks.Move(1)
			continue
		default:
			continue
		}

		break
	}

	return nodes, errors, nil
}

// seeks until anything other than space
// special cases for terminators, EOF and comments
func seekNextElem(toks *list.List[shared.Token]) (
	Result,
	shared.Token,
) {
	tok := toks.Curr()
	for ; ; tok = toks.Move(1) {
		if ttmSpacer.contains(tok.Type) {
			continue
		} else if tok.Type == shared.TOKEN_EOF {
			return RESULT_EOF, shared.Token{}
		} else if tok.Type == shared.TOKEN_COMMENT_SYM {
			return RESULT_COMMENT, shared.Token{}
		} else if ttmNodeTerminator.contains(tok.Type) {
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
