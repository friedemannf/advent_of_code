package util

import (
	"testing"
)

func TestPower(t *testing.T) {
	type args struct {
		n int
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 10,
				p: 0,
			},
			want: 1,
		}, {
			name: "",
			args: args{
				n: 10,
				p: 1,
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				n: 10,
				p: 2,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
