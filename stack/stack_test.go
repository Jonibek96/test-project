package stack

import (
	"testing"
)

func TestProcessRPN(t *testing.T) {
	type test struct {
		name   string
		input  string
		output int
	}

	tests := []test{
		{"Test  1", "3 4 +", 7},
		{"Test  2", "5 3 - 6 +", 8},
		{"Test  3", "20 5 /", 4},
		{"Test  5", "10 20 30 * +", 610},
		{"Test  6", "4 38 4 5 6 * + - /", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessRPN(tt.input)
			if got != tt.output {
				t.Errorf("processRPN() = %d, want %d", got, tt.output)
			}
		})
	}
}

func TestProcessRPNUnsuccessful(t *testing.T) {
	type test struct {
		name   string
		input  string
		output int
	}

	tests := []test{
		{"Test case 1", "20 0 /", 0},
		{"Test case 2", "20 20 - - - -", 0},
		{"Test case 3", "2 a +", 0},
		{"Test case 4", "a a / error -", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessRPN(tt.input)
			if got != tt.output {
				t.Errorf("processRPN() = %d, want %d", got, tt.output)
			}
		})
	}
}
