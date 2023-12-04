package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adjacentCells(t *testing.T) {
	type args struct {
		lines []string
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "",
			args: args{
				lines: []string{"123", "456", "789"},
				x:     0,
				y:     0,
			},
			want: []uint8{'2', '4', '5'},
		}, {
			name: "",
			args: args{
				lines: []string{"123", "456", "789"},
				x:     1,
				y:     1,
			},
			want: []uint8{'1', '2', '3', '4', '6', '7', '8', '9'},
		}, {
			name: "",
			args: args{
				lines: []string{"123", "456", "789"},
				x:     2,
				y:     1,
			},
			want: []uint8{'2', '3', '5', '8', '9'},
		}, {
			name: "",
			args: args{
				lines: []string{"123", "456", "789"},
				x:     2,
				y:     2,
			},
			want: []uint8{'5', '6', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := adjacentCells(tt.args.lines, tt.args.x, tt.args.y)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
