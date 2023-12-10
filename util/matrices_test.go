package util

import (
	"reflect"
	"testing"
)

func TestAdjacentCells(t *testing.T) {
	type args[T any] struct {
		lines [][]T
		x     int
		y     int
		oob   *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     1,
				y:     1,
			},
			want: []string{"1", "2", "3", "4", "6", "7", "8", "9"},
		}, {
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     0,
				y:     0,
				oob:   nil,
			},
			want: []string{"2", "4", "5"},
		}, {
			/*
			   X X X
			   X - 2 3
			   X 4 5 6
			     7 8 9
			*/
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     0,
				y:     0,
				oob:   New("X"),
			},
			want: []string{"X", "X", "X", "X", "2", "X", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjacentCells(tt.args.lines, tt.args.x, tt.args.y, tt.args.oob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjacentCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectingCells(t *testing.T) {
	type args[T any] struct {
		lines [][]T
		x     int
		y     int
		oob   *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     1,
				y:     1,
				oob:   nil,
			},
			want: []string{"2", "4", "6", "8"},
		}, {
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     2,
				y:     1,
				oob:   nil,
			},
			want: []string{"3", "5", "9"},
		}, {
			name: "",
			args: args[string]{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				x:     2,
				y:     2,
				oob:   New("X"),
			},
			want: []string{"6", "8", "X", "X"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectingCells(tt.args.lines, tt.args.x, tt.args.y, tt.args.oob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectingCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
