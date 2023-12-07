package main

import (
	"testing"
)

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"given example": {
			input: example,
			want:  6440,
		},
		"small case #1": {
			input: `J94T7 194
			426KQ 595
			326T9 834
			QK296 766`,
			want: 5670,
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
		want  int
	}{
		"given example": {
			input: example,
			want:  5905,
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
