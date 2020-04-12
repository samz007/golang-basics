package src

import "fmt"
import "strings"

func palindromeChecker(inputString string) bool {
	fmt.Println("Inside palindrome checker function")
	if strings.Compare(inputString, reverse_string(inputString)) == 0 {
		fmt.Println("Given input string : ", inputString, "is a palindrome")
		return true
	}
	fmt.Println("Given input string : ", inputString, "is not a palindrome")
	return false
}

func reverse_string(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
