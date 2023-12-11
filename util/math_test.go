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

func TestAbs(t *testing.T) {
	type args[T interface{ ~int }] struct {
		n T
	}
	type testCase[T interface{ ~int }] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				n: 42,
			},
			want: 42,
		}, {
			name: "",
			args: args[int]{
				n: -42,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.n); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
