package protocol

type Resource struct {
	Name       string
	Type       Type
	Class      Class
	TTL        int
	DataLength int
	Data       string
}
