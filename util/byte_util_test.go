package util

import "testing"

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
