package l75

import (
	"strings"
)

// 20231125
// LeetCode 75
// Problem 1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings
// Difficulty: Easy

// This solution can be simplified further. If str1 + str2 != str2 + str1, then return "".
// If they are equal, str[:gcd] is the greatest common divisor of str1 and str2.

// gcdOfStrings returns the greatest common divisor of strings str1 and str2.
func gcdOfStrings(str1 string, str2 string) string {
	// If str1 == str2, return str1.
	if str1 == str2 {
		return str1
	}
	// If either string is empty, return "".
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	// Find the greatest common divisor of the lengths of str1 and str2.
	gcd := getGCD(len(str1), len(str2))
	gcdStr := str1[:gcd]
	// If str1[:gcd] is a divisor of str1 and str2, return str1[:gcd].
	if str1 == strings.Repeat(gcdStr, len(str1)/gcd) && str2 == strings.Repeat(gcdStr, len(str2)/gcd) {
		return gcdStr
	}
	return ""

}

// getGCD returns the greatest common divisor of two integers.
func getGCD(a, b int) int {
	// Euclidean algorithm
	if a == 0 {
		return b
	}
	return getGCD(b%a, a)
}
