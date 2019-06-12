package util

import "../protocol"

const byteSize = 8

func ReadPacket(data []byte) {
	var offset uint = 0
	id := readInt16(offset, data)
	println("TransactionId => ", FormatHexString(id))
	offset += 2
	intFlags := readInt16(offset, data)
	var qr = protocol.QRQuery
	if intFlags>>protocol.ShiftQR&0x1 != 0 {
		qr = protocol.QRRespone
	}
	println("QR => ", qr)
	opCode := (intFlags >> protocol.ShiftOpCode) & 0xf
	println("OpCode => ", opCode)
	aa := boolBit(intFlags, protocol.ShiftAA)
	println("AA => ", aa)
	tc := boolBit(intFlags, protocol.ShiftTC)
	println("TC => ", tc)
	rd := boolBit(intFlags, protocol.ShiftRD)
	println("RD => ", rd)
	ra := boolBit(intFlags, protocol.ShiftRA)
	println("RA => ", ra)
	rCode := (intFlags >> protocol.ShiftRCode) & 0xf
	println(FormatHexString(rCode))
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
