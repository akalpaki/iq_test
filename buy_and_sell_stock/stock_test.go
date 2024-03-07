package stock

import "testing"

type testCase struct {
	name     string
	data     []int
	expected int
}

func TestFindProfitMyAlgo(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "example 2",
			data:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, tt := range tc {
		res := findProfitMyAlgo(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed for case %s: expected=%d, result=%d", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkFindProfitMyAlgo(b *testing.B) {
	bench := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		findProfitMyAlgo(bench)
	}
}

func TestFindProfitIteration1(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "example 2",
			data:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, tt := range tc {
		res := findProfitIteration1(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed for case %s: expected=%d, result=%d", tt.name, tt.expected, res)
		}
	}
}

func BenchmarkFindProfitIteration1(b *testing.B) {
	bench := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		findProfitIteration1(bench)
	}
}

func BenchmarkFindProfitIteration2(b *testing.B) {
	bench := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		findProfitIteration2(bench)
	}
}
