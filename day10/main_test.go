package main

import (
	"testing"
)

var part1Simple1 = `.....
.S-7.
.|.|.
.L-J.
.....`

var part1Simple2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var part1Complex1 = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

var part1Complex2 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"simple example #1": {
			input: part1Simple1,
			want:  4,
		},
		"simple example #2": {
			input: part1Simple2,
			want:  8,
		},
		"complex example #1": {
			input: part1Complex1,
			want:  4,
		},
		"complex example #2": {
			input: part1Complex2,
			want:  8,
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

var part2Simple = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

var part2Squeeze = `..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`

var part2Large = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

var part2Complex = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"simple example": {
			input: part2Simple,
			want:  4,
		},
		"squeeze example": {
			input: part2Squeeze,
			want:  4,
		},
		"large example": {
			input: part2Large,
			want:  8,
		},
		"complex example": {
			input: part2Complex,
			want:  10,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
