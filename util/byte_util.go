package util

import (
	"../protocol"
	"fmt"
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
		name, l := readData(offset, data)
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

	if packet.AnswerRRs > 0 {
		rs, l := readResources(offset, data, packet.AnswerRRs)
		packet.Answers = rs
		offset += l
	}

	if packet.AuthorityRRs > 0 {
		rs, l := readResources(offset, data, packet.AuthorityRRs)
		packet.Authorities = rs
		offset += l
	}

	if packet.AdditionalRRs > 0 {
		rs, l := readResources(offset, data, packet.AdditionalRRs)
		packet.Additions = rs
		offset += l
	}

	return packet
}

func readResources(offset uint, data []byte, rs int) ([]protocol.Resource, uint) {
	resources := make([]protocol.Resource, rs)
	var l uint
	for i := 0; i < rs; i++ {
		answer, length := readResource(offset, data)
		l += length
		resources[i] = answer
	}
	return resources, l
}

func readResource(offset uint, data []byte) (protocol.Resource, uint) {
	resource := protocol.Resource{}
	name, l := readName(offset, data)
	resource.Name = name
	resource.Type = protocol.ParseType(readInt16(offset+l, data))
	l += 2
	resource.Class = protocol.ParseClass(readInt16(offset+l, data))
	l += 2
	resource.TTL = readInt(offset+l, data, 32)
	l += 4
	dataLen := uint(readInt16(offset+l, data))
	l += 2

	switch resource.Type {
	case protocol.TypeA, protocol.TypeMX:
		resource.Data = fmt.Sprintf("%v.%v.%v.%v",
			readInt8(offset+l, data),
			readInt8(offset+l+1, data),
			readInt8(offset+l+2, data),
			readInt8(offset+l+3, data))
		break
	case protocol.TypeCNAME:
		name, _ := readData(offset+l, data)
		resource.Data = name
		break
	default:
	}

	l += dataLen

	return resource, l
}

func readName(offset uint, data []byte) (string, uint) {
	if data[offset] == protocol.OffsetNameFlag {
		name, _ := readData(uint(data[offset+1]), data)
		return name, 2
	}
	return readData(offset, data)
}

func readData(offset uint, data []byte) (string, uint) {
	var name string
	var i, l uint
	for data[offset+l] != 0 {
		if i > 0 {
			name += "."
		}
		strLen := uint(data[offset+l])
		l += 1
		if strLen == protocol.OffsetNameFlag {
			pos := uint(data[offset+l])
			n, ln := readData(pos, data)
			l += ln
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

func readInt8(offset uint, data []byte) int {
	return readInt(offset, data, 8)
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
