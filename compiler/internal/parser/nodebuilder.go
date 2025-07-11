package parser

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/golp/list"
)

const (
	INDENT_SPACE  IndentType = "INDENT_SPACE"
	INDENT_TAB    IndentType = "INDENT_TAB"
	INDENT_UNKOWN IndentType = "INDENT_UNKOWN"
)

type ParseCtx struct {
	Tokens        *TokenList
	StringBuilder strings.Builder
	Position      shared.Position

	IndentType IndentType
	Step       StepType
}

type RawNode struct {
	Depth    uint
	Position shared.Position

	Key              RawKey
	AssignmentSymbol string
	Value            RawValue

	CommentContent string

	Error error
}

type NodeCounter struct {
	Components  uint
	Traits      uint
	EnumMembers uint
	Comments    uint
}

type RawValue struct {
	Content    string
	ValueType  shared.TokenType
	PreSymbol  string
	PostSymbol string
}

type RawKey struct {
	Content   string
	PreSymbol string
}

type IndentType string

type TokenList = list.List[shared.Token]

func newRawNode() RawNode {
	return RawNode{}
}

func newParseCtx(toks *TokenList) ParseCtx {
	return ParseCtx{
		Tokens:        toks,
		IndentType:    INDENT_UNKOWN,
		Step:          STEP_UNDEFINED,
		StringBuilder: strings.Builder{},
	}
}

func newId(ntype shared.NodeType, counter *NodeCounter) string {
	out := ""

	switch ntype {
	case shared.NODE_TYPE_COMPONENT:
		counter.Components++
		out = fmt.Sprintf("component_%d", counter.Components)
	case shared.NODE_TYPE_TRAIT:
		counter.Traits++
		out = fmt.Sprintf("trait_%d", counter.Traits)
	case shared.NODE_TYPE_ENUM_MEMBER:
		counter.EnumMembers++
		out = fmt.Sprintf("enummember_%d", counter.EnumMembers)
	case shared.NODE_TYPE_COMMENT:
		counter.Comments++
		out = fmt.Sprintf("comment_%d", counter.Comments)
	}

	return out
}

// func (nb NodeBuilder) process(counter *NodeCounter) (shared.Node, error) {
// 	var node shared.Node
//
// 	switch nb.AssignmentSymbol {
// 	case shared.SYMBOL_TYPEDEF:
// 		node = shared.NodeComponent{
// 			Metadata: shared.NodeMetadata{
// 				Id:       newId(shared.NODE_TYPE_COMPONENT, counter),
// 				Type:     shared.NODE_TYPE_COMPONENT,
// 				Position: nb.Position,
// 				Depth:    nb.Depth,
// 			},
// 			Name:        nb.Key.Content,
// 			Type:        nb.Value.Content,
// 			IsReference: (nb.Value.PreSymbol == shared.SYMBOL_REFERENCE),
// 			IsOptional:  (nb.Value.PostSymbol == shared.SYMBOL_OPTIONAL),
// 		}
// 	// TODO: handle traits
// 	case shared.SYMBOL_DECLARE_TRAIT:
// 		node = shared.NodeTrait{
// 			Metadata: shared.NodeMetadata{
// 				Id:       newId(shared.NODE_TYPE_TRAIT, counter),
// 				Type:     shared.NODE_TYPE_TRAIT,
// 				Position: nb.Position,
// 				Depth:    nb.Depth,
// 			},
// 			Name:  nb.Key.Content,
// 			Value: nb.Value.Content,
// 		}
// 	default:
// 		if len(nb.CommentContent) > 0 {
// 			node = shared.NodeComment{
// 				Metadata: shared.NodeMetadata{
// 					Id:       newId(shared.NODE_TYPE_COMMENT, counter),
// 					Type:     shared.NODE_TYPE_COMMENT,
// 					Position: nb.Position,
// 					Depth:    nb.Depth,
// 				},
// 				Content: nb.CommentContent,
// 			}
// 		} else {
// 			node = shared.NodeEnumMember{
// 				Metadata: shared.NodeMetadata{
// 					Id:       newId(shared.NODE_TYPE_ENUM_MEMBER, counter),
// 					Type:     shared.NODE_TYPE_ENUM_MEMBER,
// 					Position: nb.Position,
// 					Depth:    nb.Depth,
// 				},
// 				Key:         nb.Key.Content,
// 				IsReference: nb.Key.PreSymbol != "",
// 				Alias:       nb.Value.Content,
// 			}
// 		}
// 	}
//
// 	return node, nil
// }
