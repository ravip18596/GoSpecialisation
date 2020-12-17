/*
Question - Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type concurrentArrays struct {
	channel1, channel2, channel3, channel4 chan *[]int
}

func main() {
	fmt.Println("Please input a series of integers which are comma-separated.")
	var intstr string
	fmt.Scan(&intstr)
	elements := strings.Split(intstr, ",")
	var arr []int
	for _, ele := range elements {
		e, err := strconv.Atoi(ele)
		if err != nil {
			fmt.Errorf("error converting input element to int. Err is %v", err)
			return
		}
		arr = append(arr, e)
	}
	n := len(arr)
	distLen := int(math.Ceil(float64(n) / 4.0))
	fmt.Println("each cell will contain this many elements - ", distLen)
	arr2 := make([][]int, 4)
	for j := 0; j < 4; j++ {
		arr2[j] = make([]int, 0)
	}
	for i, j := 0, -1; i < n && j < 4; i++ {
		if i%distLen == 0 {
			j++
		}
		arr2[j] = append(arr2[j], arr[i])
	}
	fmt.Println("sorted arr is ", launchPipeline(arr2))
}

func sortArr(arr []int, goroutine int, output chan []int) {
	fmt.Println("before sorting slice - ", arr, " goroutine is ", goroutine)
	sort.Ints(arr)
	fmt.Println("after sorting slice - ", arr, " goroutine is ", goroutine)
	output <- arr
}

func merge(arr1, arr2 []int) []int {
	var arr []int
	n1, n2 := len(arr1), len(arr2)
	var i, j int
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			arr = append(arr, arr2[j])
			j++
		}else{
			//if both equal
			arr = append(arr,arr1[i])
			arr = append(arr,arr2[j])
			i,j = i+1,j+1
		}
	}
	for i < n1 {
		arr = append(arr, arr1[i])
		i++
	}
	for j < n2 {
		arr = append(arr, arr2[j])
		j++
	}
	return arr
}

func launchPipeline(arr [][]int) []int {
	//writer section
	output := make(chan []int, 4)
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			sortArr(arr[i-1], i, output)
		}(i)
	}
	wg.Wait()
	close(output)
	//reader section
	//reading concurrently through two readers and storing the result in another reader
	wg.Add(2)
	result := make(chan []int,2)
	go func(result chan []int){
		defer wg.Done()
		arrs :=  make([][]int,0)
		for i:=0;i<2;i++{
			temp := <- output
			arrs = append(arrs,temp)
		}
		temp := merge(arrs[0],arrs[1])
		result <- temp
	}(result)
	go func(result chan []int) {
		defer wg.Done()
		arrs :=  make([][]int,0)
		for i:=0;i<2;i++{
			temp := <- output
			arrs = append(arrs,temp)
		}
		temp := merge(arrs[0],arrs[1])
		result <- temp
	}(result)
	wg.Wait()
	close(result)
	var mergedArrs [][]int
	for arr := range result{
		mergedArrs = append(mergedArrs,arr)
	}
	final := merge(mergedArrs[0],mergedArrs[1])
	fmt.Println("final - ",final)
	return final
}

func reader(output <-chan []int) []int{
	arrs := make([][]int, 2)
	for i := 0; i < 2; i++ {
		temp := <-output
		arrs[i] = make([]int,len(temp))
		//copy data from temp slice to arr slice
		copy(arrs[i],temp)
	}
	return merge(arrs[0], arrs[1])
}
