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
	flags.Z = boolBit(intFlags, protocol.ShiftZ)
	flags.AD = boolBit(intFlags, protocol.ShiftAD)
	flags.CD = boolBit(intFlags, protocol.ShiftCD)

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
		name, l := readName(offset, data, 0xff)
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

	if int(offset) >= len(data)-1 {
		return packet
	}

	answers := make([]protocol.Resource, packet.AnswerRRs)
	for i := 0; i < packet.AnswerRRs; i++ {
		answer, l := readResource(offset, data)
		offset += l
		answers[i] = answer
	}
	packet.Answers = answers

	return packet
}

func readResource(offset uint, data []byte) (protocol.Resource, uint) {
	resource := protocol.Resource{}
	var l uint
	if data[offset] == protocol.OffsetNameFlag {
		resource.Name, _ = readName(uint(data[offset+1]), data, 0xff)
		l += 2
	} else {
		name, ln := readName(offset+l, data, 0xff)
		l += ln
		resource.Name = name
	}
	println("Name =>> ", resource.Name)
	resource.Type = protocol.ParseType(readInt16(offset+l, data))
	println("Type =>> ", resource.Type.Name)
	l += 2
	resource.Class = protocol.ParseClass(readInt16(offset+l, data))
	l += 2
	resource.TTL = readInt(offset+l, data, 32)
	l += 4
	dataLen := uint(readInt16(offset+l, data))
	resource.DataLength = int(dataLen)
	l += 2
	println("DataLen =>> ", dataLen)

	if resource.Type == protocol.TypeA {
		resource.Data = string(readInt16(offset+l, data)) +
			"." + string(readInt16(offset+l+1, data)) +
			"." + string(readInt16(offset+l+2, data)) +
			"." + string(readInt16(offset+l+3, data))
	} else {
		name, _ := readName(offset+l, data, dataLen)
		resource.Data = name
	}
	l += dataLen
	println("Data =>> ", resource.Data)

	return resource, l
}

func readName(offset uint, data []byte, length uint) (string, uint) {
	var name string
	var i, l uint
	for data[offset+l] != 0 && l < length {
		if i > 0 {
			name += "."
		}
		strLen := uint(data[offset+l])
		l += 1
		if strLen == protocol.OffsetNameFlag {
			pos := uint(data[offset+l])
			n, _ := readName(pos, data, length)
			l += 1
			name += n
		} else {
			name += readString(offset+l, strLen, data)
			l += strLen
		}
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
