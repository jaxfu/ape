package shared

type (
	Ast      = []Node
	Position struct {
		Line uint
		Col  uint
	}
	NodeType = string
)

const (
	NODETYPE_COMPONENT   = "NT_COMPONENT"
	NODETYPE_COMMENT     = "NT_COMMENT"
	NODETYPE_CONSTRAINT  = "NT_CONSTRAINT"
	NODETYPE_ENUM_MEMBER = "NT_ENUM_MEMBER"
	NODETYPE_ARRAY       = "NT_ARRAY"
	NODETYPE_EMPTYLINE   = "NT_EMPTYLINE"
	NODETYPE_UNDEFINED   = "NT_UNDEFINED"
)

type Node interface {
	Meta() NodeMetadata
}

type NodeMetadata struct {
	ID       string
	Type     NodeType
	Position Position
	Depth    uint
}

type ComponentNode struct {
	Metadata      NodeMetadata
	Name          string
	ComponentType string
	IsOptional    bool
	IsReference   bool
	Children      []*Node
}

func (cn ComponentNode) Meta() NodeMetadata {
	return cn.Metadata
}

type ConstraintNode struct {
	Metadata NodeMetadata
	Name     string
	Value    string
}

func (cn ConstraintNode) Meta() NodeMetadata {
	return cn.Metadata
}

type CommentNode struct {
	Metadata NodeMetadata
	Content  string
}

func (cn CommentNode) Meta() NodeMetadata {
	return cn.Metadata
}

type EnumMemberNode struct {
	Metadata    NodeMetadata
	Key         string
	IsReference bool
	Alias       string
}

func (en EnumMemberNode) Meta() NodeMetadata {
	return en.Metadata
}
