package foo

type Blah[T any, U any] struct {
	// V1 is the first field.
	V1 T `json:"v1"`
	// V2 is the first field.
	V2 U `json:"v2"`
}
