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

// switch node.Meta().Type {
// case shared.NODETYPE_COMPONENT:
// case shared.NODETYPE_COMMENT:
// case shared.NODETYPE_CONSTRAINT:
// case shared.NODETYPE_ENUM_MEMBER:
// case shared.NODETYPE_ARRAY:
// case shared.NODETYPE_EMPTYLINE:
// case shared.NODETYPE_UNDEFINED:
// }

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

type NodeComponent struct {
	Metadata    NodeMetadata
	Name        string
	Type        string
	IsOptional  bool
	IsReference bool
}

type NodeConstraint struct {
	Metadata NodeMetadata
	Name     string
	Value    string
}

type NodeComment struct {
	Metadata NodeMetadata
	Content  string
}

type NodeEnumMember struct {
	Metadata    NodeMetadata
	Key         string
	IsReference bool
	Alias       string
}

type NodeTypesConstraint interface {
	NodeComponent |
		NodeConstraint |
		NodeComment |
		NodeEnumMember
}

// switch node.(type) {
// 	case NodeComponent:
// 	case NodeConstraint:
// 	case NodeComment:
// 	case NodeEnumMember:
// }

func (cn NodeComponent) Meta() NodeMetadata {
	return cn.Metadata
}

func (cn NodeConstraint) Meta() NodeMetadata {
	return cn.Metadata
}

func (cn NodeComment) Meta() NodeMetadata {
	return cn.Metadata
}

func (en NodeEnumMember) Meta() NodeMetadata {
	return en.Metadata
}
