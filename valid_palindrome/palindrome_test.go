package palindrome

import (
	"testing"
)

type testCase struct {
	name     string
	data     string
	expected bool
}

func TestCheckPalindromeMyAlgo(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "example 2",
			data:     "race a car",
			expected: false,
		},
		{
			name:     "example 3",
			data:     " ",
			expected: true,
		},
	}

	for _, tt := range tc {
		res := checkPalindromeMyAlgo(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed on case %s: expected=%t, result=%t", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkCheckPalindromeMyAlgo(b *testing.B) {
	bench := "A man, a plan, a canal: Panama"
	for i := 0; i < b.N; i++ {
		checkPalindromeMyAlgo(bench)
	}
}

func TestCheckPalindromeIteration1(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "example 2",
			data:     "race a car",
			expected: false,
		},
		{
			name:     "example 3",
			data:     " ",
			expected: true,
		},
	}

	for _, tt := range tc {
		res := checkPalindromeIteration1(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed on case %s: expected=%t, result=%t", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkCheckPalindromeIteration1(b *testing.B) {
	bench := "A man, a plan, a canal: Panama"
	for i := 0; i < b.N; i++ {
		checkPalindromeIteration1(bench)
	}
}
