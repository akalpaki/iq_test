package palindrome

import (
	"strings"
	"unicode"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

*/

// Naive approach: loop the string backwards, create a new string, compare the two strings
// very slow

func stripNonLetters(s string) string {
	var b strings.Builder
	s = strings.ToLower(s)
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func checkPalindromeMyAlgo(s string) bool {
	s = stripNonLetters(s)
	sLen := len(s)
	var t string
	for i := sLen - 1; i >= 0; i-- {
		t += string(s[i])
	}
	return s == t
}

// Iteration 1: I googled what two pointers technique is and I kinda remember how it works.

func checkPalindromeIteration1(s string) bool {
	s = stripNonLetters(s)
	i := 0
	j := len(s) - 1 // remember to -1 if you're indexing from the end to avoid index out of bounds
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
