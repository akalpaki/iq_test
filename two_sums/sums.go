package sums

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.
*/

// First iteration: naive approach, n^2

func twoSumMyAlgo(nums []int, target int) []int {
	for i, v := range nums {
		numsCp := make([]int, len(nums))
		copy(numsCp, nums)
		for j, k := range numsCp {
			if i == j {
				continue // we can only use a number once.
			}
			if v+k == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumIteration1(nums []int, target int) []int {
	return []int{}
}
