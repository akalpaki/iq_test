package main

import (
	"slices"
)

/*
Given a list of rocks' weights and an integer maxVariance, divide the rocks into groups such that
within a group the difference in weight between any two rocks is less than or equal to maxVariance.
Find the minimum number of groups to divide all the rocks.

Example:

weights = [5,3,2,4,6]
maxVariance = 2
Output: 2
Explanation: divide into two groups [3,2,4] and [5,6]
*/

/*
My approach:
Step 1: sort the array
Step 2: Loop through the array and group elements together. Groups contain the indexes of the elements.

	Step 2.1: Grouping
	- Start the group by adding the index of the current element.
	- Loop through the array. If rocks[j] - rocks[i] <= maxVariance, add j's index to the group;
	else break this loop and continue to next i.
	- Store the group on a map[int][]int

Step 3: Take the first group and loop through the map:
  - For each number in the group (dst), check if the other group (src) has the same number. If yes, continue;
    else append src to dst and continue.
  - Each time you append, add 1 to a counter.

Step 4: Return the counter
*/

/*
	IMPORTANT NOTE:
	When you iterate over the same slice twice in a nested loop, before the second iteration you have to
	copy the original slice and iterate over the copy. That is because slices are pointers over an under-
	lying array, so if you use the same pointer at both rounds of the iteration, each round will move it along
	and eventually it will reach the end and get stuck there, leaving you with a bug.
*/

func check_shared_idx(base []int32, other []int32) bool {
	for _, idxB := range base {
		for _, idxO := range other {
			if idxB == idxO {
				return true
			}
		}
	}
	return false
}

func Group_division(data []int32, variance int32) int32 {
	// Step 1
	slices.Sort[[]int32, int32](data)
	groups := make(map[int][]int32)
	// Step 2
	dataLen := len(data) // avoid uneccessary calls to len(group)
	for i := 0; i < dataLen; i++ {
		group := make([]int32, 0)
		dataCp := make([]int32, dataLen) // copy the group array so we have a fresh array to iterate over
		copy(dataCp, data)
		group = append(group, int32(i))
		if i == dataLen-1 {
			groups[i] = group
			break
		}
		for j := i + 1; j <= dataLen; j++ {
			if j == dataLen {
				groups[i] = group
				break
			}
			if dataCp[j]-data[i] <= variance {
				group = append(group, int32(j))
			} else {
				groups[i] = group
				break
			}
		}
	}
	// Step 3
	counter := int32(1)
	base := groups[0]
	for _, group := range groups {
		if check_shared_idx(base, group) {
			continue
		}
		counter++
		base = append(base, group...)
	}
	return counter
}

func GDiv2(data []int, numRange int) int {
	dataLen := len(data)
	if dataLen == 0 {
		return 0
	}
	slices.Sort(data)
	groupCount := 1
	var maxVal int
	for i, v := range data {
		if i == 0 {
			maxVal = v + numRange
			continue
		}

		if i == dataLen {
			groupCount++
			break
		}

		if v > maxVal {
			maxVal = v + numRange
			groupCount++
		}
	}

	return groupCount
}
