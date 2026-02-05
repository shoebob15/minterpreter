package main

import (
	"testing"
)

func TestInterpreterEval(t *testing.T) {
	var tests = []struct {
		input       string
		expected    int
		shouldError bool
	}{
		{"2+3", 5, false},
		{"4 + 6", 10, false},
		{"12 + 25", 37, false},
		{"2-3", -1, false},
		{"4 - 6", -2, false},
		{"12 - 25", -13, false},
		{"-", 0, true},
		{":", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			interpreter, err := NewInterpreter(tt.input)
			if err != nil && !tt.shouldError {
				t.Error("failed test: could not create interpreter")
			}

			result, err := interpreter.Eval()
			if result != tt.expected || (err != nil && !tt.shouldError) {
				t.Errorf("failed test: expected %d, got %d; err: %s",
					tt.expected,
					result,
					err)
			}
		})
	}
}
