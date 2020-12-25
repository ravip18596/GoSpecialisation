/*
Problem - Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop,
the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int, 0, 3)
	var input string
	for {
		fmt.Println("Input number below. Input x to stop.")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}
		if strings.EqualFold(input, "x") {
			break
		}
		in, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("err - converting input string. err is ", err)
			continue
		}
		arr = append(arr, in)
		sort.Ints(arr)
		fmt.Println("sorted slice is ", arr)
	}
}
