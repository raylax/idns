package util

import "testing"

func TestFormatHexString(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: struct{ v int }{v: 15},
			want: "0x000f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatHexString(tt.args.v); got != tt.want {
				t.Errorf("FormatHexString() = %v, want %v", got, tt.want)
			}
		})
	}
}
