package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdjacentCells(t *testing.T) {
	type args[T comparable] struct {
		matrix Matrix[T]
		coord  Coordinate
		oob    *T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "",
			args: args[string]{
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 1,
					Y: 1,
				},
			},
			want: []string{"1", "2", "3", "4", "6", "7", "8", "9"},
		}, {
			name: "",
			args: args[string]{
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 0,
					Y: 0,
				},
				oob: nil,
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
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 0,
					Y: 0,
				},
				oob: New("X"),
			},
			want: []string{"X", "X", "X", "X", "2", "X", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.matrix.AdjacentCells(tt.args.coord, tt.args.oob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjacentCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectingCells(t *testing.T) {
	type args[T comparable] struct {
		matrix Matrix[T]
		coord  Coordinate
		oob    *T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "",
			args: args[string]{
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 1,
					Y: 1,
				},
				oob: nil,
			},
			want: []string{"2", "4", "6", "8"},
		}, {
			name: "",
			args: args[string]{
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 2,
					Y: 1,
				},
				oob: nil,
			},
			want: []string{"3", "5", "9"},
		}, {
			name: "",
			args: args[string]{
				matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
				coord: Coordinate{
					X: 2,
					Y: 2,
				},
				oob: New("X"),
			},
			want: []string{"6", "8", "X", "X"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.matrix.ConnectingCells(tt.args.coord, tt.args.oob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectingCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	type args[T comparable] struct {
		matrix Matrix[T]
		dir    Modifier
	}
	type testCase[T comparable] struct {
		name    string
		args    args[T]
		wantOut [][]T
	}
	tests := []testCase[int]{
		{
			name: "horizontal forward",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				dir: Horizontal | Forward,
			},
			wantOut: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		}, {
			name: "horizontal reverse",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				dir: Horizontal | Reverse,
			},
			wantOut: [][]int{
				{3, 2, 1},
				{6, 5, 4},
				{9, 8, 7},
			},
		}, {
			name: "vertical forward",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: Vertical | Forward,
			},
			wantOut: [][]int{
				{1, 4, 7, 10},
				{2, 5, 8, 11},
				{3, 6, 9, 12},
			},
		}, {
			name: "vertical reverse",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: Vertical | Reverse,
			},
			wantOut: [][]int{
				{10, 7, 4, 1},
				{11, 8, 5, 2},
				{12, 9, 6, 3},
			},
		}, {
			name: "diagonal LR forward",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: DiagonalLR | Forward,
			},
			wantOut: [][]int{
				{10},
				{7, 11},
				{4, 8, 12},
				{1, 5, 9},
				{2, 6},
				{3},
			},
		}, {
			name: "diagonal LR reverse",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: DiagonalLR | Reverse,
			},
			wantOut: [][]int{
				{10},
				{11, 7},
				{12, 8, 4},
				{9, 5, 1},
				{6, 2},
				{3},
			},
		}, {
			name: "diagonal RL forward",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: DiagonalRL | Forward,
			},
			wantOut: [][]int{
				{12},
				{9, 11},
				{6, 8, 10},
				{3, 5, 7},
				{2, 4},
				{1},
			},
		}, {
			name: "diagonal RL reverse",
			args: args[int]{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				dir: DiagonalRL | Reverse,
			},
			wantOut: [][]int{
				{12},
				{11, 9},
				{10, 8, 6},
				{7, 5, 3},
				{4, 2},
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.args.matrix.Transform(tt.args.dir); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Transform() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestMatrix_FloodFill(t *testing.T) {
	type args struct {
		coordinate Coordinate
	}
	type testCase[T comparable] struct {
		name   string
		matrix Matrix[T]
		args   args
		want   []Coordinate
	}
	tests := []testCase[int]{
		{
			name: "",
			matrix: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			args: args{coordinate: Coordinate{
				X: 1,
				Y: 2,
			}},
			want: []Coordinate{
				{1, 0},
				{2, 0},
				{1, 1},
				{2, 1},
				{1, 2},
			},
		}, {
			name: "",
			matrix: [][]int{
				{0, 1, 1, 1},
				{0, 1, 3, 1},
				{2, 1, 1, 1},
				{1, 2, 2, 0},
			},
			args: args{coordinate: Coordinate{
				X: 2,
				Y: 0,
			}},
			want: []Coordinate{
				{1, 0},
				{2, 0},
				{3, 0},
				{1, 1},
				{3, 1},
				{1, 2},
				{2, 2},
				{3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.FloodFill(tt.args.coordinate)
			require.ElementsMatch(t, tt.want, got, "FloodFill() = %v, want %v", got, tt.want)
		})
	}
}

func TestMatrix_Height(t *testing.T) {
	type testCase[T comparable] struct {
		name      string
		m         Matrix[T]
		want      int
		wantPanic bool
	}
	tests := []testCase[int]{
		{
			name: "",
			m: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 2,
		},
		{
			name: "",
			m: [][]int{
				{1, 2},
				{3},
			},
			want:      0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected a panic")
					}
				}()
			}
			if got := tt.m.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
