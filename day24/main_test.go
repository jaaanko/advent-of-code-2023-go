package main

import (
	"testing"
)

var example = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		upper int
		lower int
		want  int
	}{
		"given example": {
			input: example,
			upper: 27,
			lower: 7,
			want:  2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tt.input, tt.upper, tt.lower)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
