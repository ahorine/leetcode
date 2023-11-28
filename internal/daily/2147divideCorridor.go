package daily

// 20231128
// Problem 2147: Number of Ways to Divide a Long Corridor
// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor
// Difficulty: Hard

// numberOfWays returns the number of ways to divide the corridor.
func numberOfWays(corridor string) int {
	var (
		ans, seenSeats, seenPlants int
	)
	// Iterate over each character in the corridor.
	for _, c := range corridor {
		if c == 'S' {
			seenSeats++
			if seenSeats == 2 && ans < 1 {
				// We only start caring about the number of plants when we've seen two seats.
				ans = 1
			} else if seenSeats == 3 {
				// We've seen a third seat, so we sum the answer with seenPlants+1, with the modulo to prevent an overflow.
				ans = ans * (seenPlants + 1) % (1e9 + 7)
				// Reset the seen seats and plants.
				seenSeats = 1
				seenPlants = 0
			}
		} else if seenSeats == 2 {
			// We only care about the number of plants when we've seen two seats.
			seenPlants++
		}
	}
	// If there are an odd number of seats, then there are no ways to divide the corridor.
	if seenSeats == 1 {
		return 0
	}
	return ans
}
