package internal

import (
	"fmt"
)

type Ast struct {
	RootNodes []*DeclarationNode `json:"rootNodes"`
}

// represents Root component node
type DeclarationNode struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	Constraints []*ConstraintNode  `json:"constraints"`
	Depth       int                `json:"depth"`
	Parent      *DeclarationNode   `json:"-"` // omit to prevent cycles
	Children    []*DeclarationNode `json:"children"`
	// only for node.Type == enum
	EnumKeys []EnumKey `json:"enum_keys"`
}

// only used as a child of dec node
type ConstraintNode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EnumKey = struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p *Parser) newDeclarationNode() *DeclarationNode {
	p.DeclarationCount += 1
	return &DeclarationNode{
		Id: fmt.Sprintf(
			"declaration_%d",
			p.DeclarationCount,
		),
		Constraints: []*ConstraintNode{},
		Children:    []*DeclarationNode{},
		EnumKeys:    []EnumKey{},
	}
}

func NewParser() Parser {
	return Parser{
		Tokens: []Token{},
	}
}

type Parser struct {
	Tokens           []Token
	Cursor           int
	DeclarationCount int
}

func (p *Parser) Parse(tokens []Token) (Ast, error) {
	p.setup(tokens)

	// entry: dec name
	if p.curr().Type != toktypes.COMPONENT_NAME {
		return Ast{}, fmt.Errorf(
			"invalid syntax line %d: invalid token '%s'",
			p.curr().LineNum,
			p.curr().Value,
		)
	}

	ast := Ast{
		RootNodes: []*DeclarationNode{},
	}
	// main recursive func
	for p.curr().Type != toktypes.EOF {
		switch p.curr().Type {
		case toktypes.COMPONENT_NAME:
			node, err := p.parseDecNode(0)
			if err != nil {
				return Ast{}, fmt.Errorf("Parser.parseDecNode: %+v", err)
			}
			// fmt.Printf("%+v\n", node)
			ast.RootNodes = append(ast.RootNodes, node)

		default:
			p.move(1)
		}
	}

	return ast, nil
}

func (p *Parser) parseDecNode(
	depth int,
) (
	*DeclarationNode,
	error,
) {
	// basic sancheck
	if p.curr().Type != toktypes.COMPONENT_NAME {
		return nil, fmt.Errorf("expected declaration name")
	}

	// parse without checking prefixes on first iter
	node := p.newDeclarationNode()
	node.Depth = depth
	if err := p.parseDeclarationMetadata(node); err != nil {
		return nil, fmt.Errorf("parsing error: %+v", err)
	}

	for iter := 1; p.curr().Type != toktypes.EOF; iter++ {
		switch p.curr().Type {
		case toktypes.COMPONENT_NAME, toktypes.COMMENT_SYMBOL:
			childPrfxCnt := 0
			if p.curr().Type == toktypes.COMMENT_SYMBOL {
				childPrfxCnt += 1
				for i := 1; p.look(i).Type == toktypes.COMMENT_SYMBOL; i++ {
					childPrfxCnt += 1
				}
			}

			if (childPrfxCnt == depth || childPrfxCnt == 0) && iter != 0 {
				return node, nil
			} else if childPrfxCnt < depth {
				return node, nil
			} else if childPrfxCnt > depth {
				p.move(childPrfxCnt)
				child, err := p.parseDecNode(childPrfxCnt)
				if err != nil {
					return nil, err
				}
				node.Children = append(node.Children, child)
				child.Parent = node
			} else {
				return nil, fmt.Errorf("unhandled outcome")
			}

		case toktypes.CONSTRAINT_NAME:
			cons, err := p.parseConstraintNode()
			if err != nil {
				return nil, fmt.Errorf("parsing error: %+v", err)
			}
			node.Constraints = append(node.Constraints, cons)

		case toktypes.ENUM_KEY:
			ek := EnumKey{
				Key: p.curr().Value,
			}
			if p.look(1).Type == toktypes.ENUM_VALUE {
				ek.Value = p.move(1).Value
			} else {
				ek.Value = ""
			}
			node.EnumKeys = append(node.EnumKeys, ek)
			p.move(1)

		default:
			p.move(1)
		}
	}

	return node, nil
}

func (p *Parser) parseDeclarationMetadata(node *DeclarationNode) error {
	node.Name = p.curr().Value

	if p.move(1).Type != toktypes.COMPONENT_DEFINITION_SYMBOL {
		return fmt.Errorf(
			"error line %d:\nexpected '%s' after name '%s'",
			p.Tokens[p.Cursor].LineNum,
			SYMBOL_DEFINE_COMPONENT,
			node.Name,
		)
	}
	if p.move(1).Type != toktypes.COMPONENT_TYPE &&
		p.curr().Type != toktypes.COMPONENT_TYPE_ENUM {
		return fmt.Errorf(
			"error line %d:\nexpected value in declaration for '%s'",
			p.Tokens[p.Cursor].LineNum,
			node.Name,
		)
	}
	node.Type = p.curr().Value

	p.move(1)
	return nil
}

func (p *Parser) parseConstraintNode() (*ConstraintNode, error) {
	cons := &ConstraintNode{
		Key:   "",
		Value: "",
	}
	cons.Key = p.curr().Value

	if p.move(1).Type != toktypes.CONSTRAINT_DEFINITION_SYMBOL {
		return nil, fmt.Errorf(
			"error line %d:\nexpected '%s' after key '%s'",
			p.Tokens[p.Cursor].LineNum,
			SYMBOL_DEFINE_CONSTRAINT,
			cons.Key,
		)
	}
	if p.move(1).Type != toktypes.CONSTRAINT_VALUE {
		return nil, fmt.Errorf(
			"error line %d:\nexpected value in declaration for '%s'",
			p.curr().LineNum,
			cons.Key,
		)
	}
	cons.Value = p.curr().Value

	p.move(1)
	return cons, nil
}

// return val at cursor
func (p *Parser) curr() Token {
	return p.Tokens[p.Cursor]
}

// look without updating cursor +/-,
// overflow returns start or end
func (p *Parser) look(jmps int) Token {
	toklen := len(p.Tokens)

	if jmps < 0 && (p.Cursor+jmps < 0) {
		return p.Tokens[0]
	} else if jmps > 0 && p.Cursor+jmps >= toklen {
		return p.Tokens[toklen-1]
	}

	return p.Tokens[p.Cursor+jmps]
}

// update cursor +/-, return val,
// overflow returns start or end
func (p *Parser) move(jmps int) Token {
	toklen := len(p.Tokens)
	index := 0

	if jmps < 0 && (p.Cursor+jmps < 0) {
		index = 0
	} else if jmps > 0 && p.Cursor+jmps >= toklen {
		index = toklen - 1
	} else {
		index = p.Cursor + jmps
	}

	p.Cursor = index
	return p.curr()
}

func (p *Parser) setup(tokens []Token) {
	p.Tokens = tokens
	p.Cursor = 0
	p.DeclarationCount = 0
}
