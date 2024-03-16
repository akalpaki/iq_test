package cycle

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can
be reached again by continuously following the next pointer. Internally, pos is used to denote the index of
the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type Node struct {
	Val  int
	Next *Node
}

func walk(head *Node, visited map[*Node]struct{}) bool {
	if head == nil {
		return false
	}

	if _, ok := visited[head]; ok {
		return true
	}

	visited[head] = struct{}{}
	head = head.Next
	return walk(head, visited)
}

// first approach: recursive solution
func hasCycle(head *Node) bool {
	visited := make(map[*Node]struct{})
	return walk(head, visited)
}
