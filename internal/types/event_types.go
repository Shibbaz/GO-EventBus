package types

type EventArgs map[string]any

type Projection struct{}

type Status struct{
	Code int
	Message string
}