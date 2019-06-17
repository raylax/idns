package protocol

type Resource struct {
	Name  string
	Type  Type
	Class Class
	TTL   int
	Data  string
}
