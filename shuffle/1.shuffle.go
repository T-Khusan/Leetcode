package main

import "fmt"

// https://leetcode.com/problems/shuffle-string/
// You are given a string s
// and an integer array indices of the same length.
// The string s will be shuffled
// such that the character at the ith position moves to indices[i] in the shuffled string.
// Return the shuffled string.

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices))
	for i, v := range indices {
		result[v] = s[i]
	}
	return string(result)
}

func main() {
	jim := restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})
	fmt.Println(jim)
}
