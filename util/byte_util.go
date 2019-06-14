package util

import (
	"../protocol"
)

const byteSize = 8

func ReadPacket(data []byte) protocol.Protocol {
	packet := protocol.Protocol{}
	var offset uint = 0

	id := readInt16(offset, data)
	packet.TransactionId = id
	offset += 2

	intFlags := readInt16(offset, data)
	flags := protocol.Flags{}

	var qr = protocol.QRQuery
	if intFlags>>protocol.ShiftQR&0x1 != 0 {
		qr = protocol.QRResponse
	}
	flags.QR = qr

	opCode := (intFlags >> protocol.ShiftOpCode) & 0xf
	flags.OpCode = protocol.ParseOpCode(opCode)

	flags.AA = boolBit(intFlags, protocol.ShiftAA)
	flags.TC = boolBit(intFlags, protocol.ShiftTC)
	flags.RD = boolBit(intFlags, protocol.ShiftRD)
	flags.RA = boolBit(intFlags, protocol.ShiftRA)

	flags.RCode = protocol.ParseRCode((intFlags >> protocol.ShiftRCode) & 0xf)
	offset += 2

	packet.Flags = flags

	packet.QuestionRRs = readInt16(offset, data)
	offset += 2

	packet.AnswerRRs = readInt16(offset, data)
	offset += 2

	packet.AuthorityRRs = readInt16(offset, data)
	offset += 2

	packet.AdditionalRRs = readInt16(offset, data)
	offset += 2

	questions := make([]protocol.Question, packet.QuestionRRs)
	for i := 0; i < packet.QuestionRRs; i++ {
		name, l := readDomain(offset, data)
		offset += l
		ntype := protocol.ParseType(readInt16(offset, data))
		offset += 2
		nclass := protocol.ParseClass(readInt16(offset, data))
		offset += 2
		question := protocol.Question{
			Name:  name,
			Type:  ntype,
			Class: nclass,
		}
		questions[i] = question
	}
	packet.Questions = questions

	return packet
}

func readDomain(offset uint, data []byte) (string, uint) {
	var name string
	var i, l uint
	for data[offset+l] != 0 {
		if i > 0 {
			name += "."
		}
		length := uint(data[offset+l])
		l += 1
		name += readString(offset+l, length, data)
		l += length
		i++
	}
	return name, l + 1
}

func boolBit(flags int, shift uint) bool {
	return (flags>>shift)&0x1 == 1
}

func readInt16(offset uint, data []byte) int {
	return readInt(offset, data, 16)
}

func readInt(offset uint, data []byte, bit int) int {
	var r = 0
	bytes := uint(bit / byteSize)
	for c := bytes; c > 0; c-- {
		r |= int(data[offset+c-1]) << ((bytes - c) * byteSize)
	}
	return r
}

func readString(offset uint, length uint, data []byte) string {
	return string(data[offset : offset+length])
}
