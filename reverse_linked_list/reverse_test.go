package reverse

import "testing"

type testCase struct {
	name     string
	data     *ListNode
	expected []int
}

func TestReverseListMyAlgo(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "example 2",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expected: []int{2, 1},
		},
		{
			name:     "example 3",
			data:     nil,
			expected: nil,
		},
	}

	for _, tt := range tc {
		res := reverseListMyAlgo(tt.data)
		if tt.expected == nil {
			if res != nil {
				t.Errorf("Test failed at case %s: expected=%v, result=%p", tt.name, tt.expected, res)
				t.FailNow()
			}
			continue
		}
		var compare []int
		current := res
		for current.Next != nil {
			compare = append(compare, current.Val)
			current = current.Next
		}
		compare = append(compare, current.Val)

		for i, v := range compare {
			if v != tt.expected[i] {
				t.Errorf("Test failed at case %s: expected=%v, result=%v", tt.name, tt.expected, compare)
			}
		}
	}
}

func BenchmarkReverseListMyAlgo(b *testing.B) {
	bench := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	for i := 0; i < b.N; i++ {
		reverseListMyAlgo(bench)
	}
}

func TestReverseListIteration1(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "example 2",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expected: []int{2, 1},
		},
		{
			name:     "example 3",
			data:     nil,
			expected: nil,
		},
	}

	for _, tt := range tc {
		res := reverseListIteration1(tt.data)
		if tt.expected == nil {
			if res != nil {
				t.Errorf("Test failed at case %s: expected=%v, result=%p", tt.name, tt.expected, res)
				t.FailNow()
			}
			continue
		}
		var compare []int
		current := res
		for current.Next != nil {
			compare = append(compare, current.Val)
			current = current.Next
		}
		compare = append(compare, current.Val)

		for i, v := range compare {
			if v != tt.expected[i] {
				t.Errorf("Test failed at case %s: expected=%v, result=%v", tt.name, tt.expected, compare)
			}
		}
	}
}

func BenchmarkReverseListIteration1(b *testing.B) {
	bench := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	for i := 0; i < b.N; i++ {
		reverseListIteration1(bench)
	}
}

func TestReverseListIteration2(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "example 2",
			data:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expected: []int{2, 1},
		},
		{
			name:     "example 3",
			data:     nil,
			expected: nil,
		},
		{
			name:     "example 4",
			data:     &ListNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tc {
		res := reverseListIteration2(tt.data)
		if tt.expected == nil {
			if res != nil {
				t.Errorf("Test failed at case %s: expected=%v, result=%p", tt.name, tt.expected, res)
				t.FailNow()
			}
			continue
		}
		var compare []int
		current := res
		for current.Next != nil {
			compare = append(compare, current.Val)
			current = current.Next
		}
		compare = append(compare, current.Val)

		for i, v := range compare {
			if v != tt.expected[i] {
				t.Errorf("Test failed at case %s: expected=%v, result=%v", tt.name, tt.expected, compare)
			}
		}
	}
}

func BenchmarkReverseListIteration2(b *testing.B) {
	bench := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	for i := 0; i < b.N; i++ {
		reverseListIteration2(bench)
	}
}
