package main

import (
	"testing"
)

var example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"given example": {
			input: example,
			want:  136,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tt.input)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		spins int
		want  int
	}{
		"given example": {
			input: example,
			spins: 1000000000,
			want:  64,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tt.input, tt.spins)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
