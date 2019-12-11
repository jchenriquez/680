package main

import "fmt"

func validPalindromeH(s string, i, j int, jumpsTaken int) bool {
	if jumpsTaken >= 2 {
		return false
	}

	if i >= j {
		return true
	}

	if s[i] == s[j] {
		return validPalindromeH(s, i+1, j-1, jumpsTaken)
	}

	return validPalindromeH(s, i+1, j, jumpsTaken + 1) || validPalindromeH(s, i, j-1, jumpsTaken+1)
}

func validPalindrome(s string) bool {
	return validPalindromeH(s, 0, len(s)-1, 0)
}

func main() {
	fmt.Printf("palindrome is valid %v\n", validPalindrome("aba"))
}
