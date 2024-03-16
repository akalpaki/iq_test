package invert

import "testing"

type testCase struct {
	name     string
	data     *TreeNode
	expected []int
}

func bfsToArray(root *TreeNode) []int {
	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current == nil {
			continue
		}

		q = append(q, current.Left)
		q = append(q, current.Right)
		res = append(res, current.Val)
	}
	return res
}

func TestInvertTree(t *testing.T) {
	tc := []testCase{
		{
			name:     "Example 1",
			data:     &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2",
			data:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Example 3",
			data:     nil,
			expected: []int{},
		},
	}

	for _, tt := range tc {
		res := invertTree(tt.data)
		resArray := bfsToArray(res)
		if len(resArray) == 0 && len(tt.expected) != 0 {
			t.Errorf("Test failed for case %s: expected=%v, res=%v", tt.name, tt.expected, resArray)
			t.FailNow()
		}
		for i := 0; i < len(resArray); i++ {
			if tt.expected[i] != resArray[i] {
				t.Errorf("Test failed for case %s: expected=%v, res=%v", tt.name, tt.expected, resArray)
				t.FailNow()
			}
		}
	}
}
