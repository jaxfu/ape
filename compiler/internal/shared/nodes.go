package shared

const (
	NODE_TYPE_COMPONENT   = "COMPONENT"
	NODE_TYPE_COMMENT     = "COMMENT"
	NODE_TYPE_TRAIT       = "TRAIT"
	NODE_TYPE_ENUM_MEMBER = "ENUM_MEMBER"
	NODE_TYPE_ARRAY       = "ARRAY"
	NODE_TYPE_EMPTYLINE   = "EMPTYLINE"
	NODE_TYPE_UNDEFINED   = "UNDEFINED"
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

type NodeTrait struct {
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
		NodeTrait |
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

func (cn NodeTrait) Meta() NodeMetadata {
	return cn.Metadata
}

func (cn NodeComment) Meta() NodeMetadata {
	return cn.Metadata
}

func (en NodeEnumMember) Meta() NodeMetadata {
	return en.Metadata
}
