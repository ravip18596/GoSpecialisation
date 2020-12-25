/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that
the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(arr []int, i int) {
	//swap arr[i] with arr[i+1]
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j)
			}
		}
	}
}

func main() {
	fmt.Println("Please enter sequence of up to 10 integers in a single line separated by , ")
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0, 10)
	if scanner.Scan() {
		elementStr := scanner.Text()
		elements := strings.Split(elementStr, ",")
		for _, ele := range elements {
			eleInt, err := strconv.Atoi(ele)
			if err != nil {
				fmt.Println("err is ", err)
				continue
			}
			arr = append(arr, eleInt)
		}
		BubbleSort(arr)
		result := strings.Builder{}
		for _, ele := range arr {
			result.WriteString(strconv.Itoa(ele))
			result.WriteByte(' ')
		}
		fmt.Println("sorted array is ", result.String())
	}
}
