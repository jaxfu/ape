package parser

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/shared"
)

const (
	STEP_COMMENT         StepType = "COMMENT"
	STEP_ENTRY           StepType = "ENTRY"
	STEP_DEPTH           StepType = "DEPTH"
	STEP_KEY             StepType = "KEY"
	STEP_ASSIGNMENT_SYM  StepType = "ASSIGNMENT_SYM"
	STEP_VALUE           StepType = "VALUE"
	STEP_SEEK_TERMINATOR StepType = "SEEK_TERMINATOR"
	STEP_UNDEFINED       StepType = "UNDEFINED"
)

type StepType string

type Step interface {
	stype() StepType
	process(ParseCtx, RawNode) (ParseCtx, RawNode, Step, error)
}

// entry
type StepEntry struct {
	stepType StepType
}

var stepEntry = StepEntry{stepType: STEP_ENTRY}

func (st StepEntry) stype() StepType {
	return st.stepType
}

func (st StepEntry) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	tok := ctx.Tokens.Curr()
	tt := tok.Type
	if ttmNodeEntry.contains(tt) {
		node.Position = tok.Position
		return ctx, node, stepKey, nil
	} else if ttmCommentEntry.contains(tt) {
		node.AssignmentSymbol = tok.Content
		node.Position = tok.Position
		return ctx, node, stepComment, nil
	} else if ttmIndent.contains(tt) {
		node.Position = tok.Position
		return ctx, node, stepDepth, nil
	} else {
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected token '%s'\n", tok.Content))
	}
}

// comment
type StepComment struct {
	stepType StepType
}

var stepComment = StepComment{stepType: STEP_COMMENT}

func (st StepComment) stype() StepType {
	return st.stepType
}

func (st StepComment) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	tok := ctx.Tokens.Curr()
	tt := tok.Type

	if ttmNodeTerminator.contains(tt) {
		node.CommentContent = ctx.StringBuilder.String()
		return ctx, node, nil, nil
	} else {
		ctx.StringBuilder.WriteString(tok.Content)
		ctx.Tokens.Move(1)
		return ctx, node, stepComment, nil
	}
}

// key
type StepKey struct {
	stepType StepType
}

var stepKey = StepKey{stepType: STEP_KEY}

func (st StepKey) stype() StepType {
	return st.stepType
}

func (st StepKey) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	tok := ctx.Tokens.Curr()
	// prefix symbol check
	if tok.Type == shared.TOKEN_SYMBOL {
		if tok.Content == shared.SYMBOL_OPTIONAL {
			node.Key.PreSymbol = shared.SYMBOL_OPTIONAL
			ctx.Tokens.Move(1)
		} else if tok.Content == shared.SYMBOL_REFERENCE {
			// is enum if key starts with ref sym
			node.Key.PreSymbol = tok.Content
			ctx.Tokens.Move(1)
		} else {
			return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected symbol %s", tok.Content))
		}
	}

	if tok.Type != shared.TOKEN_IDENT {
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, "unexpected token")
	}
	node.Key.Content = tok.Content

	ctx.Tokens.Move(1)
	return ctx, node, stepAssignmentSym, nil
}

// assignment sym
type StepAssignmentSym struct {
	stepType StepType
}

var stepAssignmentSym = StepAssignmentSym{stepType: STEP_ASSIGNMENT_SYM}

func (st StepAssignmentSym) stype() StepType {
	return st.stepType
}

func (st StepAssignmentSym) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	tok := ctx.Tokens.Curr()

	res, tok := seekTokType(
		ctx.Tokens,
		shared.TOKEN_SYMBOL,
	)
	switch res {
	case RESULT_TERMINATOR:
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, "unexpected node ending")
	case RESULT_EOF:
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, "unexpected EOF")
	case RESULT_UNEXPECTED:
		if ttmValue.contains(tok.Type) {
			return ctx, node, stepValue, nil
		} else {
			return ctx, node, nil, shared.NewSyntaxError(tok.Position, "unexpected character")
		}
	}

	node.AssignmentSymbol = parseSymbol(ctx.Tokens)
	ctx.Tokens.Move(1)
	return ctx, node, stepValue, nil
}

// value
type StepValue struct {
	stepType StepType
}

var stepValue = StepValue{stepType: STEP_VALUE}

func (st StepValue) stype() StepType {
	return st.stepType
}

func (st StepValue) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	toks := ctx.Tokens
	res, tok := ttmValue.seekMember(toks)
	switch res {
	case RESULT_TERMINATOR:
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected token '%s'\n", tok.Content))
	case RESULT_UNEXPECTED:
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected token '%s'\n", tok.Content))
	case RESULT_EOF:
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected token '%s'\n", tok.Content))
	}

	// check for prefix sym
	if tok.Type == shared.TOKEN_SYMBOL {
		node.Value.PreSymbol = parseSymbol(toks)
		tok = toks.Curr()
	}

	// get value content
	if !ttmValue.contains(tok.Type) {
		return ctx, node, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected token '%s'\n", tok.Content))
	} else {
		node.Value.Content = tok.Content
		tok = toks.Move(1)
	}

	// check for postfix sym
	if tok.Type == shared.TOKEN_SYMBOL {
		node.Value.PostSymbol = parseSymbol(toks)
	}

	return ctx, node, nil, nil
}

// depth
type StepDepth struct {
	stepType StepType
}

var stepDepth = StepDepth{stepType: STEP_DEPTH}

func (st StepDepth) stype() StepType {
	return st.stepType
}

func (st StepDepth) process(
	ctx ParseCtx,
	node RawNode,
) (
	ParseCtx,
	RawNode,
	Step,
	error,
) {
	toks := ctx.Tokens
	tok := toks.Curr()

	if ctx.IndentType == INDENT_UNKOWN { // indentation has not been set
		if tok.Type == shared.TOKEN_TAB {
			ctx.IndentType = INDENT_TAB
		} else if tok.Type == shared.TOKEN_SPACE {
			ctx.IndentType = INDENT_SPACE
		}
	} else { // ensure consistent indentation
		if ctx.IndentType == INDENT_SPACE &&
			tok.Type == shared.TOKEN_TAB {
			return ctx, node, nil, shared.NewSyntaxError(tok.Position, "inconsistent indentation, expected space")
		} else if ctx.IndentType == INDENT_TAB &&
			tok.Type == shared.TOKEN_SPACE {
			return ctx, node, nil, shared.NewSyntaxError(tok.Position, "inconsistent indentation, expected tab")
		}
	}

	node.Depth++
	toks.Move(1)
	return ctx, node, stepEntry, nil
}
