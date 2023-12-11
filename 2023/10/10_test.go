package day10

import (
	"testing"
)

func Test_shoelaceArea(t *testing.T) {
	type args struct {
		path [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				path: [][]int{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
			},
			want: 4,
		}, {
			name: "",
			args: args{
				path: [][]int{{1, 6}, {3, 1}, {7, 2}, {4, 4}, {8, 5}},
			},
			want: 16.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shoelaceArea(tt.args.path); got != tt.want {
				t.Errorf("shoelaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
