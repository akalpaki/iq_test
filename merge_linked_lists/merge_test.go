package merge

import (
	"testing"
)

type testCase struct {
	name     string
	input1   *Node
	input2   *Node
	expected []int
}

func TestMergeLinkedListsMyAlgo(t *testing.T) {
	tc := []testCase{
		{
			name:     "example 1",
			input1:   &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 4}}},
			input2:   &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 4}}},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "example 2",
			input1:   nil,
			input2:   nil,
			expected: nil,
		},
		{
			name:     "example 3",
			input1:   nil,
			input2:   &Node{Val: 0},
			expected: []int{0},
		},
		{
			name:     "example 4",
			input1:   &Node{Val: 2},
			input2:   &Node{Val: 1},
			expected: []int{1, 2},
		},
	}
	for _, tt := range tc {
		head := mergeLinkedListsMyAlgo(tt.input1, tt.input2)
		if head == nil && tt.expected == nil {
			continue
		}
		res := []int{}
		current := head
		for current.Next != nil {
			res = append(res, current.Val)
			current = current.Next
		}
		res = append(res, current.Val)
		if len(res) != len(tt.expected) {
			t.Errorf("Test failed for case %s: expected and result not same length", tt.name)
			t.Errorf("Expected: %v, Result %v", tt.expected, res)
			t.FailNow()
		}
		for i, v := range res {
			if v != tt.expected[i] {
				t.Errorf("Test failed for case %s: expected=%v, result=%v", tt.name, tt.expected, res)
				t.FailNow()
			}
		}
	}
}

func BenchmarkMergeLinkedListsMyAlgo(b *testing.B) {
	input1 := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 4}}}
	input2 := &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 4}}}

	for i := 0; i < b.N; i++ {
		mergeLinkedListsMyAlgo(input1, input2)
	}
}
