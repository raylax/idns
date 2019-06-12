package protocol

type Resource struct {
	Name

	Type

	Class

	TTL

	DataLength int16

	Data string
}
