package parser

type RawApeObject struct {
	Name        *string           `json:"name"`
	Category    *string           `json:"category"`
	Description *string           `json:"description"`
	Props       RawApeObjectProps `json:"props"`
}

type RawApeObjectProps map[string]map[string]any
