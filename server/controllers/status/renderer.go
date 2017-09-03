package status

type Renderer interface {
	String(view string, data interface{}) (string, error)
}
