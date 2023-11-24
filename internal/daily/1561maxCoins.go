package daily

import "slices"

// 20231124
// Problem 1561: Maximum Number of Coins You Can Get
// https://leetcode.com/problems/maximum-number-of-coins-you-can-get
// Difficulty: Medium

// maxCoins returns the maximum number of coins you can get by following the rules below:
// - There are 3n piles of coins of varying size.
// - In each step, you will choose any 3 piles of coins (not necessarily consecutive).
// - Take the second largest pile of coins and add it to your collection.
// - Return the maximum number of coins you can get.
func maxCoins(piles []int) int {
	var ans int
	// Sort the piles in ascending order.
	slices.Sort(piles)
	// Drop the first n/3 piles, as we'll never choose them.
	piles = piles[len(piles)/3:]
	// Sum every other pile, as we'll always choose the largest pile.
	for i := 0; i < len(piles); i += 2 {
		ans += piles[i]
	}
	return ans
}
