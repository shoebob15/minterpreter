package main

import (
	"testing"
)

func TestInterpreterEval(t *testing.T) {
	var tests = []struct {
		input       string
		expected    float64
		shouldError bool
	}{
		{"2+3", 5, false},        // addition
		{"4 + 6", 10, false},     // whitespace
		{"12 + 25", 37, false},   // abritrary length integers
		{"4 - 6", -2, false},     // subraction
		{"12 - 25", -13, false},  // negative result
		{"5 * 6", 30, false},     // multiplication
		{"20 / 5", 4, false},     // division
		{"18 * 2 -5", 31, false}, // pemdas
		{"5 + 20*4", 85, false},  // more pemdas
		{"5 * (7-2)", 25, false}, // parenthesis
		{"5 += 0", 0, true},      // error handling
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			interpreter, err := NewInterpreter(tt.input)
			if err != nil && !tt.shouldError {
				t.Error("failed test: could not create interpreter")
			}

			result, err := interpreter.Eval()
			if result != tt.expected || (err != nil && !tt.shouldError) {
				t.Errorf("failed test: expected %v, got %v; err: %s",
					tt.expected,
					result,
					err)
			}
		})
	}
}
