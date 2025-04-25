package components

type ReferenceTag = string

type Reference struct {
	TargetId ComponentId
	IsLinked bool
	Target   *Component
}
