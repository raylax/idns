package util

const byteSize = 8

func MakePacket(data []byte) {
	var offset uint = 0
	id := readInt(offset, data, 16)
	println(FormatHexString(id))
	offset += 2
	//flagsInt := readInt(offset, data, 16)
}

func readInt(offset uint, data []byte, bit int) int {
	var r = 0
	bytes := uint(bit / byteSize)
	for c := bytes; c > 0; c-- {
		r |= int(data[offset+c-1]) << ((bytes - c) * byteSize)
	}
	return r
}
