package protocol

type Resource struct {
	Name       Name
	Type       Type
	Class      Class
	TTL        int
	DataLength int16
	Data       string
}
