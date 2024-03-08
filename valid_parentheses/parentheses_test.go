package parentheses

import "testing"

func TestLifo(t *testing.T) {
	lifo := NewLifo()
	lifo.Push('{')
	if v, ok := lifo.GetPrevious(); !ok || v != '{' {
		t.Errorf("Expected 1, got %d", v)
	}
	lifo.Pull()
	if v, ok := lifo.GetPrevious(); ok || v != byte(0) {
		t.Errorf("Expected 0, got %d", v)
	}
}

type testCase struct {
	name     string
	data     string
	expected bool
}

func TestIsValidMyAlgo(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     "()",
			expected: true,
		},
		{
			name:     "example 2",
			data:     "()[]{}",
			expected: true,
		},
		{
			name:     "example 3",
			data:     "(]",
			expected: false,
		},
		{
			name:     "empty string",
			data:     "",
			expected: false,
		},
		{
			name:     "ignore non-brackets",
			data:     "(a)c[ ]_{}",
			expected: true,
		},
		{
			name:     "nested brackets",
			data:     "({})",
			expected: true,
		},
		{
			name:     "nested error",
			data:     "({[]}",
			expected: false,
		},
	}

	for _, tt := range tc {
		res := isValidMyAlgo(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed on case %s: expected=%t, result=%t", tt.name, tt.expected, res)
		}
	}
}
