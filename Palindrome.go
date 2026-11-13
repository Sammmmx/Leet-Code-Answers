package main

import "fmt"

func longestPalindrome(s string) string {
	longest := ""

	for i := 0; i < len(s); i++ {

		palindrome1 := expand(s, i, i)
		if len(palindrome1) > len(longest) {
			longest = palindrome1
		}

		palindrome2 := expand(s, i, i+1)
		if len(palindrome2) > len(longest) {
			longest = palindrome2
		}
	}
	return longest
}

func expand(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func main() {
	result := longestPalindrome("babad")
	fmt.Println(result)
}
