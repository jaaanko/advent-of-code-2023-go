package main

import (
	"testing"
)

var example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		gap   int
		want  int
	}{
		"given example": {
			input: example,
			gap:   2,
			want:  374,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tt.input, tt.gap)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		gap   int
		want  int
	}{
		"given example with gap of 10": {
			input: example,
			gap:   10,
			want:  1030,
		},
		"given example with gap of 100": {
			input: example,
			gap:   100,
			want:  8410,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tt.input, tt.gap)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
