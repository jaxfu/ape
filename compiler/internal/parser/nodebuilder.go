package parser

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
)

const (
	PREALLOC uint = 1024

	INDENT_SPACE  IndentType = "INDENT_SPACE"
	INDENT_TAB    IndentType = "INDENT_TAB"
	INDENT_UNKOWN IndentType = "INDENT_UNKOWN"
)

type NodeBuilder struct {
	NodeType   shared.NodeType
	IndentType IndentType

	StringBuilder *strings.Builder

	Depth    uint
	Position shared.Position

	Key      RawKey
	Assigner string
	Value    RawValue

	CommentContent string

	Error error
}

type NodeCounter struct {
	Components  uint
	Constraints uint
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

func NewNodeBuilder(indType IndentType) NodeBuilder {
	return NodeBuilder{
		NodeType:      shared.NODETYPE_UNDEFINED,
		IndentType:    indType,
		StringBuilder: &strings.Builder{},
	}
}

func newId(ntype shared.NodeType, counter *NodeCounter) string {
	out := ""

	switch ntype {
	case shared.NODETYPE_COMPONENT:
		counter.Components++
		out = fmt.Sprintf("component_%d", counter.Components)
	case shared.NODETYPE_CONSTRAINT:
		counter.Constraints++
		out = fmt.Sprintf("constraint_%d", counter.Constraints)
	case shared.NODETYPE_ENUM_MEMBER:
		counter.EnumMembers++
		out = fmt.Sprintf("enummember_%d", counter.EnumMembers)
	case shared.NODETYPE_COMMENT:
		counter.Comments++
		out = fmt.Sprintf("comment_%d", counter.Comments)
	}

	return out
}

func (nb NodeBuilder) process(counter *NodeCounter) (shared.Node, error) {
	var node shared.Node

	switch nb.Assigner {
	case shared.SYMBOL_DECLARE_COMPONENT:
		node = shared.ComponentNode{
			Metadata: shared.NodeMetadata{
				Id:       newId(shared.NODETYPE_COMPONENT, counter),
				Type:     shared.NODETYPE_COMPONENT,
				Position: nb.Position,
				Depth:    nb.Depth,
			},
			Name:          nb.Key.Content,
			ComponentType: nb.Value.Content,
			IsReference:   (nb.Value.PreSymbol == shared.SYMBOL_MARK_REFERENCE),
			IsOptional:    (nb.Value.PostSymbol == shared.SYMBOL_MARK_OPTIONAL),
		}
	case shared.SYMBOL_DECLARE_CONSTRAINT:
		node = shared.ConstraintNode{
			Metadata: shared.NodeMetadata{
				Id:       newId(shared.NODETYPE_CONSTRAINT, counter),
				Type:     shared.NODETYPE_CONSTRAINT,
				Position: nb.Position,
				Depth:    nb.Depth,
			},
			Name:  nb.Key.Content,
			Value: nb.Value.Content,
		}
	default:
		if len(nb.CommentContent) > 0 {
			node = shared.CommentNode{
				Metadata: shared.NodeMetadata{
					Id:       newId(shared.NODETYPE_COMMENT, counter),
					Type:     shared.NODETYPE_COMMENT,
					Position: nb.Position,
					Depth:    nb.Depth,
				},
				Content: nb.CommentContent,
			}
		} else {
			node = shared.EnumMemberNode{
				Metadata: shared.NodeMetadata{
					Id:       newId(shared.NODETYPE_ENUM_MEMBER, counter),
					Type:     shared.NODETYPE_ENUM_MEMBER,
					Position: nb.Position,
					Depth:    nb.Depth,
				},
				Key:         nb.Key.Content,
				IsReference: nb.Key.PreSymbol != "",
				Alias:       nb.Value.Content,
			}
		}
	}

	return node, nil
}
