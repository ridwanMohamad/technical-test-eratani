package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Input kata palindrom : ")
	fmt.Scan(&str)
	if isPalindrome(str) {
		fmt.Printf("%s itu palindrom.", str)
	} else {
		fmt.Printf("%s itu bukan palindrome.", str)
	}
}
func isPalindrome(input string) bool {
	// Remove spaces and convert the string to lowercase
	input = strings.Trim(input, " ")

	// Compare the string with its reverse
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}

	return true
}
