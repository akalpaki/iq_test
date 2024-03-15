package merge

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

type Node struct {
	Val  int
	Next *Node
}

func mergeLinkedListsMyAlgo(list1, list2 *Node) *Node {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list2 != nil && list1 == nil {
		return list2
	}

	out := &Node{}
	outTail := out

	curr1 := list1
	curr2 := list2

	// first comparison happens out of loop to give out correct initial value
	if curr1.Val >= curr2.Val {
		out.Val = curr2.Val
		curr2 = curr2.Next
	} else {
		out.Val = curr1.Val
		curr1 = curr1.Next
	}

	for curr1 != nil && curr2 != nil {
		if curr1.Val >= curr2.Val {
			node := &Node{Val: curr2.Val}
			outTail.Next = node
			outTail = node
			curr2 = curr2.Next
		} else {
			node := &Node{Val: curr1.Val}
			outTail.Next = node
			outTail = node
			curr1 = curr1.Next
		}
	}

	if curr2 == nil && curr1 != nil {
		outTail.Next = curr1
	}
	if curr1 == nil && curr2 != nil {
		outTail.Next = curr2
	}

	return out
}
