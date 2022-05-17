package verr

type Context map[string]any

func NewContext(v map[string]any) Context {
	return v
}
