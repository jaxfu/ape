package schemas

type Prop struct {
	Name         string
	Type         string
	IntFields    *PropTypeIntegerFields
	StringFields *PropTypeStringFields
}

type PropTypeIntegerFields struct {
	Size uint `json:"size"`
}

type PropTypeStringFields struct {
	MinLength uint `json:"min_length"`
	MaxLength uint `json:"max_length"`
}
