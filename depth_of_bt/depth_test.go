package depth

import "testing"

/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

type testCase struct {
	name     string
	data     *TreeNode
	expected int
}

func TestMaxDepth(t *testing.T) {
	tc := []testCase{
		{
			name:     "Example 1",
			data:     &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expected: 3,
		},
		{
			name:     "Example 2",
			data:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: 2,
		},
		{
			name:     "Example 3",
			data:     nil,
			expected: 0,
		},
	}

	for _, tt := range tc {
		res := maxDepth(tt.data)
		if res != tt.expected {
			t.Errorf("Test failed for case %s: expected=%d, result=%d", tt.name, tt.expected, res)
			t.FailNow()
		}
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	bench := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	for i := 0; i < b.N; i++ {
		maxDepth(bench)
	}
}
