package shared

const (
	NODETYPE_COMPONENT   = "NT_COMPONENT"
	NODETYPE_COMMENT     = "NT_COMMENT"
	NODETYPE_CONSTRAINT  = "NT_CONSTRAINT"
	NODETYPE_ENUM_MEMBER = "NT_ENUM_MEMBER"
	NODETYPE_ARRAY       = "NT_ARRAY"
	NODETYPE_EMPTYLINE   = "NT_EMPTYLINE"
	NODETYPE_UNDEFINED   = "NT_UNDEFINED"
)

type Ast = []Node

type Position struct {
	Line uint
	Col  uint
}

type NodeType = string

type Node interface {
	Meta() NodeMetadata
}

type NodeMetadata struct {
	Id       string
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
}

type ConstraintNode struct {
	Metadata NodeMetadata
	Name     string
	Value    string
}

type CommentNode struct {
	Metadata NodeMetadata
	Content  string
}

type EnumMemberNode struct {
	Metadata    NodeMetadata
	Key         string
	IsReference bool
	Alias       string
}

func (cn ComponentNode) Meta() NodeMetadata {
	return cn.Metadata
}

func (cn ConstraintNode) Meta() NodeMetadata {
	return cn.Metadata
}

func (cn CommentNode) Meta() NodeMetadata {
	return cn.Metadata
}

func (en EnumMemberNode) Meta() NodeMetadata {
	return en.Metadata
}
