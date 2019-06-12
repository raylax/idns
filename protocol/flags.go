package protocol

type Flags struct {
	QR
	OpCode
	AA F
	TC F
	RD F
	RA F
	Z  int
	RCode
}

type F bool

type QR int8

const (
	QRQuery   QR = 0
	QRRespone QR = 1
)

type OpCode int

const (
	OpCodequery  OpCode = 0x0000
	OpCodeiquery OpCode = 0x0001
	OpCodestatus OpCode = 0x0002
	OpCodenotify OpCode = 0x0004
	OpCodeupdate OpCode = 0x0005
)

type RCode int8

const (
	RCodeNoerror  RCode = 0
	RCodeServfail RCode = 2
	RCodeNxdomain RCode = 3
)
