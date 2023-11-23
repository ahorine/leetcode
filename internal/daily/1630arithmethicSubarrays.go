package daily

import "slices"

// 20231123
// Problem 1630: Arithmetic Subarrays
// https://leetcode.com/problems/arithmetic-subarrays
// Difficulty: Medium

// This solution is not the most efficient, as there's a number theory solution that doesn't
// require sorting. However, this solution is simple and easy to understand, and still passes
// all the test cases in LeetCode.

// checkArithmeticSubarrays returns a slice of booleans indicating whether each subarray
// defined by the given l and r slices is an arithmetic sequence.
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var results []bool
	for i := 0; i < len(l); i++ {
		// We need to make a copy of the subarray because we're going to sort it.
		subarray := make([]int, r[i]-l[i]+1)
		copy(subarray, nums[l[i]:r[i]+1])
		results = append(results, isArithmetic(subarray))
	}
	return results
}

// isArithmetic returns true if the given slice of integers is an arithmetic sequence, and
// false otherwise.
func isArithmetic(nums []int) bool {
	// Base cases that don't require sorting.
	if len(nums) < 2 {
		return false
	} else if len(nums) == 2 {
		return true
	}
	// Sort the slice and check if the difference between each pair of numbers is the same.
	slices.Sort(nums)
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}
