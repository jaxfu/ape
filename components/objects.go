package components

type Object struct {
	ComponentMetadata
	Props PropsMap `json:"props,omitempty" toml:"props,omitempty"`
}
