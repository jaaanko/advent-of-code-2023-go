package main

import (
	"testing"
)

var example = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		steps int
		want  int
	}{
		"given example": {
			input: example,
			steps: 6,
			want:  16,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tt.input, tt.steps)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
