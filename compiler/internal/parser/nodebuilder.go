package parser

import (
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
)

const (
	PREALLOC uint = 1024
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

const (
	INDENT_SPACE  IndentType = "INDENT_SPACE"
	INDENT_TAB    IndentType = "INDENT_TAB"
	INDENT_UNKOWN IndentType = "INDENT_UNKOWN"
)

type IndentType string

func NewNodeBuilder(indType IndentType) NodeBuilder {
	return NodeBuilder{
		NodeType:      shared.NODETYPE_UNDEFINED,
		IndentType:    indType,
		StringBuilder: &strings.Builder{},
	}
}

func (nb NodeBuilder) cast() (shared.Node, error) {
	var node shared.Node

	switch nb.Assigner {
	case shared.SYMBOL_DECLARE_COMPONENT:
		node = shared.ComponentNode{
			Metadata: shared.NodeMetadata{
				ID:       "TestComponent",
				Type:     shared.NODETYPE_COMPONENT,
				Position: nb.Position,
				Depth:    nb.Depth,
			},
			Name:          nb.Key.Content,
			ComponentType: nb.Value.Content,
			IsReference:   (nb.Value.PreSymbol == shared.SYMBOL_MARK_REFERENCE),
			IsOptional:    (nb.Value.PostSymbol == shared.SYMBOL_MARK_OPTIONAL),
			Children:      []*shared.Node{},
		}
	case shared.SYMBOL_DECLARE_CONSTRAINT:
		node = shared.ConstraintNode{
			Metadata: shared.NodeMetadata{
				ID:       "TestConstraint",
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
					ID:       "TestComment",
					Type:     shared.NODETYPE_COMMENT,
					Position: nb.Position,
					Depth:    nb.Depth,
				},
				Content: nb.CommentContent,
			}
		} else {
			node = shared.EnumMemberNode{
				Metadata: shared.NodeMetadata{
					ID:       "TestEnumMember",
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
