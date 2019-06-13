package util

import (
	"testing"
)

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

func Test_boolBit(t *testing.T) {
	type args struct {
		flags int
		shift uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: struct {
				flags int
				shift uint
			}{flags: 0x000f, shift: 1},
			want: true,
		},
		{
			args: struct {
				flags int
				shift uint
			}{flags: 0x000f, shift: 3},
			want: true,
		},
		{
			args: struct {
				flags int
				shift uint
			}{flags: 0x000f, shift: 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boolBit(tt.args.flags, tt.args.shift); got != tt.want {
				t.Errorf("boolBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readString(t *testing.T) {
	type args struct {
		offset uint
		length uint
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
				length uint
				data   []byte
			}{offset: 1, length: 3, data: []byte("\x03\x77\x77\x77\x05\x62\x61\x69\x64\x75\x03\x63\x6f\x6d\x00")},
			want: "www",
		},
		{
			args: struct {
				offset uint
				length uint
				data   []byte
			}{offset: 0, length: 3, data: []byte("\x77\x77\x77\x05\x62\x61\x69\x64\x75\x03\x63\x6f\x6d\x00")},
			want: "www",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readString(tt.args.offset, tt.args.length, tt.args.data); got != tt.want {
				t.Errorf("readString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInt16(t *testing.T) {
	type args struct {
		offset uint
		data   []byte
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
			}{offset: 0, data: []byte("\xf\xf")},
			want: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInt16(tt.args.offset, tt.args.data); got != tt.want {
				t.Errorf("readInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInt(t *testing.T) {
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
			}{offset: 0, data: []byte("\xf\xf"), bit: 8},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInt(tt.args.offset, tt.args.data, tt.args.bit); got != tt.want {
				t.Errorf("readInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
