package parser

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/list"
	"github.com/jaxfu/ape/compiler/internal/shared"
)

const (
	STEP_ENTRY           = "STEP_ENTRY"
	STEP_DEPTH           = "STEP_DEPTH"
	STEP_KEY             = "STEP_KEY"
	STEP_ASSIGNER        = "STEP_ASSIGNER"
	STEP_SEEK_TERMINATOR = "STEP_SEEK_TERMINATOR"
	STEP_VALUE           = "STEP_VALUE"
	STEP_COMMENT         = "STEP_COMMENT"
)

type StepType = string

type Step interface {
	stype() StepType
	process(*list.List[shared.Token], NodeBuilder) (NodeBuilder, Step, error)
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
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	tok := toks.Curr()
	tt := tok.Type
	if isIn(tt, nodeEntryMap) {
		nb.Position = tok.Position
		return nb, stepKey, nil
	} else if isIn(tt, commentEntryMap) {
		nb.Assigner = tok.Content
		nb.Position = tok.Position
		return nb, stepComment, nil
	} else if isIn(tt, indentMap) {
		nb.Position = tok.Position
		return nb, stepDepth, nil
	} else {
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected character")
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
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	tok := toks.Curr()
	tt := tok.Type

	if isIn(tt, nodeTerminatorMap) {
		nb.CommentContent = nb.StringBuilder.String()
		return nb, nil, nil
	}

	nb.StringBuilder.WriteString(tok.Content)
	toks.Move(1)
	return nb, stepComment, nil
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
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	tok := toks.Curr()
	// if key starts with symbol,
	if tok.Type == shared.TOKEN_SYMBOL {
		// check if ref symbol
		if tok.Content != shared.SYMBOL_MARK_REFERENCE {
			return nb, nil, shared.NewSyntaxError(tok.Position, fmt.Sprintf("unexpected symbol %s", tok.Content))
		}
		// if so, is enum member
		nb.NodeType = shared.NODETYPE_ENUM_MEMBER
		nb.Key.PreSymbol = tok.Content
		toks.Move(1)
	}

	if tok.Type != shared.TOKEN_IDENT {
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected token")
	}
	nb.Key.Content = tok.Content

	toks.Move(1)
	return nb, stepAssigner, nil
}

// assigner
type StepAssigner struct {
	stepType StepType
}

var stepAssigner = StepAssigner{stepType: STEP_ASSIGNER}

func (st StepAssigner) stype() StepType {
	return st.stepType
}

func (st StepAssigner) process(
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	tok := toks.Curr()
	if isConstraint(tok) {
		nb.Assigner = shared.SYMBOL_DECLARE_CONSTRAINT
		nb.NodeType = shared.NODETYPE_CONSTRAINT
		toks.Move(1)
		return nb, stepValue, nil
	}

	res, tok := seekTokType(toks, shared.TOKEN_SYMBOL)
	switch res {
	case RESULT_TERMINATOR:
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected node ending")
	case RESULT_EOF:
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected EOF")
	case RESULT_UNEXPECTED:
		// if value type, is enum member,
		// end step and return stepValue
		if isIn(tok.Type, valueMap) {
			nb.NodeType = shared.NODETYPE_ENUM_MEMBER
			return nb, stepValue, nil
		} else {
			return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected character")
		}
	}

	asn := parseSymbol(toks)

	switch asn {
	case shared.SYMBOL_DECLARE_COMPONENT:
		nb.NodeType = shared.NODETYPE_COMPONENT
		nb.Assigner = asn
	}

	toks.Move(1)
	return nb, stepValue, nil
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
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	res, tok := seekValueEntry(toks)
	switch res {
	case RESULT_TERMINATOR:
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected node ending")
	case RESULT_UNEXPECTED:
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected character")
	case RESULT_EOF:
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected EOF")
	}

	// get pre-symbol (if any)
	if tok.Type == shared.TOKEN_SYMBOL {
		if tok.Content != shared.SYMBOL_MARK_REFERENCE {
			return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected symbol")
		} else {
			nb.Value.PreSymbol = shared.SYMBOL_MARK_REFERENCE
			tok = toks.Move(1)
		}
	}
	// get value content
	if !isIn(tok.Type, valueMap) {
		return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected token")
	} else {
		nb.Value.Content = tok.Content
		tok = toks.Move(1)
	}
	// get post-symbol (if any)
	if tok.Type == shared.TOKEN_SYMBOL {
		if tok.Content != shared.SYMBOL_MARK_OPTIONAL {
			return nb, nil, shared.NewSyntaxError(tok.Position, "unexpected symbol")
		} else {
			nb.Value.PostSymbol = shared.SYMBOL_MARK_OPTIONAL
			toks.Move(1)
		}
	}

	return nb, nil, nil
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
	toks *list.List[shared.Token],
	nb NodeBuilder,
) (
	NodeBuilder,
	Step,
	error,
) {
	tok := toks.Curr()

	if nb.IndentType == INDENT_UNKOWN { // indentation has not been set
		if tok.Type == shared.TOKEN_TAB {
			nb.IndentType = INDENT_TAB
		} else if tok.Type == shared.TOKEN_SPACE {
			nb.IndentType = INDENT_SPACE
		}
	} else { // ensure consistent indentation
		if nb.IndentType == INDENT_SPACE &&
			tok.Type == shared.TOKEN_TAB {
			return nb, nil, shared.NewSyntaxError(tok.Position, "inconsistent indentation, expected space")
		} else if nb.IndentType == INDENT_TAB &&
			tok.Type == shared.TOKEN_SPACE {
			return nb, nil, shared.NewSyntaxError(tok.Position, "inconsistent indentation, expected tab")
		}
	}

	nb.Depth++
	toks.Move(1)
	return nb, stepEntry, nil
}
