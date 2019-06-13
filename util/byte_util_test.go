package util

import (
	"testing"
)

func TestReadInt(t *testing.T) {
	type args struct {
		offset uint
		data   []byte
		bit    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: struct {
				offset uint
				data   []byte
				bit    int
			}{offset: 0, data: []byte{byte(99)}, bit: 8},
			want: 99,
		},
		{
			args: struct {
				offset uint
				data   []byte
				bit    int
			}{offset: 1, data: []byte{byte(0), byte(99)}, bit: 8},
			want: 99,
		},
		{
			args: struct {
				offset uint
				data   []byte
				bit    int
			}{offset: 0, data: []byte{byte(12), byte(34)}, bit: 16},
			want: 3106,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInt(tt.args.offset, tt.args.data, tt.args.bit); got != tt.want {
				t.Errorf("ReadInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDomain(t *testing.T) {
	type args struct {
		offset uint
		data   []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: struct {
				offset uint
				data   []byte
			}{offset: 0, data: []byte("\x03\x77\x77\x77\x05\x62\x61\x69\x64\x75\x03\x63\x6f\x6d\x00")},
			want: "www.baidu.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := readDomain(tt.args.offset, tt.args.data); got != tt.want {
				t.Errorf("readDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
