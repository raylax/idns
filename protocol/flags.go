package protocol

type Flags struct {
	QR
	OpCode
	AA		F
	TC		F
	RD		F
	RA		F
	Z		int8
	RCode
}

type F bool

type QR int8

const (
	QR_QUERY	QR = 0
	QR_RESPONE	QR = 1
)

type OpCode int8

const (
	OPCODE_QUERY	OpCode = 0x0000
	OPCODE_IQUERY	OpCode = 0x0001
	OPCODE_STATUS	OpCode = 0x0002
	OPCODE_NOTIFY	OpCode = 0x0004
	OPCODE_UPDATE	OpCode = 0x0005
)

type RCode int8

const (
	RCODE_NOERROR	RCode = 0
	RCODE_SERVFAIL	RCode = 2
	RCODE_NXDOMAIN	RCode = 3
)