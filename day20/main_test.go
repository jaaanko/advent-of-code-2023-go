package main

import (
	_ "embed"
	"testing"
)

var example1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var example2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func Test_part1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"given example #1": {
			input: example1,
			want:  32000000,
		},
		"given example #2": {
			input: example2,
			want:  11687500,
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

//go:embed actual.txt
var actual string

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"actual": {
			input: actual,
			want:  224046542165867,
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
