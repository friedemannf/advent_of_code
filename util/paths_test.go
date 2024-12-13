package util

import (
	"testing"
)

func Test_ShoelaceArea(t *testing.T) {
	type args struct {
		path [][2]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				path: [][2]int{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
			},
			want: 4,
		}, {
			name: "",
			args: args{
				path: [][2]int{{1, 6}, {3, 1}, {7, 2}, {4, 4}, {8, 5}},
			},
			want: 16.5,
		}, {
			name: "",
			args: args{
				path: [][2]int{{0, 0}, {0, 2}, {2, 2}, {2, 3}, {3, 3}, {3, 0}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShoelaceArea(tt.args.path); got != tt.want {
				t.Errorf("ShoelaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
