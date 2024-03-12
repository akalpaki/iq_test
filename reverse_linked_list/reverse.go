package reverse

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

A linked list can be traversed either iteratively or recursively. Implement both ways.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListMyAlgo(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	values := []int{}
	current := head
	for current.Next != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	values = append(values, current.Val) // get the last value from the list

	out := &ListNode{Val: values[len(values)-1]}
	current = out
	for i := len(values) - 2; i >= 0; i-- {
		temp := &ListNode{Val: values[i]}
		current.Next = temp
		current = current.Next
	}

	return out
}

// Iteration 1: remove values array, directly build the reversed linked list

func reverseListIteration1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//edge case: 1 item list
	if head.Next == nil {
		return head
	}

	out := &ListNode{Val: head.Val, Next: nil} // rather than store results, add them to the head of the output linked list
	current := head.Next
	// edge case: linked list has two items, so current.Next is nil
	if current.Next == nil {
		temp := &ListNode{Val: current.Val, Next: out}
		out = temp
		return out
	}
	for current.Next != nil {
		temp := &ListNode{Val: current.Val, Next: out}
		out = temp
		current = current.Next
		if current.Next == nil {
			last := &ListNode{Val: current.Val, Next: out}
			out = last
		}
	}
	return out
}

// Iteration 2: recursive solution

func reverseListIteration2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	out := &ListNode{Val: head.Val}
	if head.Next == nil { // edge case: 1 node linked list
		return out
	}
	current := head.Next
	// construct closure for recursion
	var flip func(h *ListNode) *ListNode
	flip = func(h *ListNode) *ListNode {
		if h.Next == nil {
			temp := &ListNode{Val: h.Val, Next: out}
			out = temp
			return out
		} else {
			temp := &ListNode{Val: h.Val, Next: out}
			out = temp
			return flip(h.Next)
		}
	}
	return flip(current)
}
