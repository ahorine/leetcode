package l75

// 20231124
// LeetCode 75
// Problem 1768: Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately
// Difficulty: Easy

// Optimal solution evidently uses a string builder.

// mergeAlternately returns the merge of word1 and word2, where the merge is formed by
// alternating the letters in word1 and word2, starting with word1. If one of the words
// is longer than the other, then the remaining letters in the longer word should be
// appended at the end of the merge.
func mergeAlternately(word1 string, word2 string) string {
	var ans []byte
	// Iterate over the shorter word, then append the remaining letters of the longer word.
	short := min(len(word1), len(word2))
	for i := 0; i < short; i++ {
		ans = append(ans, word1[i], word2[i])
	}

	// Append the remaining letters of the longer word.
	if len(word1) > len(word2) {
		ans = append(ans, word1[short:]...)
	} else if len(word2) > len(word1) {
		ans = append(ans, word2[short:]...)
	}

	return string(ans)
}
