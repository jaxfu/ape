package internal

type AST struct {
	Nodes []*DeclarationNode
}

type DeclarationNode struct {
	Name        string
	Type        string
	Constraints []ConstraintNode
	Children    []*DeclarationNode
}

func NewDeclarationNode() *DeclarationNode {
	return &DeclarationNode{
		Constraints: []ConstraintNode{},
		Children:    []*DeclarationNode{},
	}
}

type ConstraintNode struct {
	Key   string
	Value string
}

func Parse([]Token) (AST, error) {
	return AST{}, nil
}
