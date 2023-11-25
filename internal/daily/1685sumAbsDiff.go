package daily

// 20231125
// Problem 1685: Sum of Absolute Differences in a Sorted Array
// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array
// Difficulty: Medium

// Turns out calculating the suffix sum wasn't necessary, and likely cost some memory.
// Per a comment on the problem, just the prefix is needed:
// result[i] = (nums[i]*i - pre[i]) + (pre[n]-pre[i] - nums[i]*(n-i))

// getSumAbsoluteDifferences returns an array ans of length n, where ans[i] is the sum of
// absolute differences between nums[i] and all the other elements in the array.
func getSumAbsoluteDifferences(nums []int) []int {
	var (
		ans, prefix, suffix = make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
		n                   = len(nums)
	)
	// Calculate the prefix and suffix sums.
	prefix[0] = nums[0]
	suffix[n-1] = nums[n-1]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
		suffix[n-i-1] = suffix[n-i] + nums[n-i-1]
	}
	// We can calculate the answer for each index i by using the prefix and suffix sums.
	for i := 0; i < len(nums); i++ {
		left := i*nums[i] - prefix[i]
		right := suffix[i] - nums[i]*(n-i-1)
		ans[i] = left + right
	}
	return ans
}
