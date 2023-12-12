package main

import (
	"testing"
)

var example = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"given example": {
			input: example,
			want:  21,
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
			want:  525152,
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
