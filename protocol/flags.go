package protocol

type Flags struct {
	QR     QR
	OpCode OpCode
	AA     bool
	TC     bool
	RD     bool
	RA     bool
	Z      bool
	AD     bool
	CD     bool
	RCode  RCode
}
