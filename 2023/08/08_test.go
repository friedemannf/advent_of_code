package day8

import (
	"testing"
)

func TestLCM(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				ints: []int{2, 3},
			},
			want: 6,
		}, {
			name: "",
			args: args{
				ints: []int{2, 3, 4},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.ints...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
