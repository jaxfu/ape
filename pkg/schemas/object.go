package schemas

type Object struct {
	ObjectId    string  `json:"object_id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description *string `json:"description,omitempty"`
	Props       []Prop
}
