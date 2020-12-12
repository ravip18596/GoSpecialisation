package main

import (
	"fmt"
	"os"
	"strconv"
)

func convertFloatToInt(floatStr string) int {
	floatNo, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Errorf("incorrect floating number provided. Err is %v", err)
		return -1
	}
	var fint int
	fint = int(floatNo)
	return fint
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Please enter floating number in command line.")
		return
	}
	fmt.Printf("input floating point number is %s \n", args[0])
	fmt.Println(convertFloatToInt(args[0]))
}
