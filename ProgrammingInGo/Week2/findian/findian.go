package main

import (
	"fmt"
	"strings"
)

func check(in string) bool {
	lower := strings.ToLower(in)
	strRune := []rune(lower)
	if strRune[0] == 'i' && strRune[len(strRune)-1] == 'n' && strings.ContainsRune(lower, 'a') {
		return true
	}
	return false
}

func main() {
	fmt.Println("Enter a string after the prompt starting with i, containing a and ending with n")
	var input string
	fmt.Scan(&input)
	if check(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
