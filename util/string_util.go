package util

import "strconv"

func FormatHexString(v int) string {
	hex := strconv.FormatInt(int64(v), 16)
	for len(hex) < 4 {
		hex = "0" + hex
	}
	return "0x" + hex
}
