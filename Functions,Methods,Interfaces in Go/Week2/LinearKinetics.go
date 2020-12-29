/*
Problem Statement -
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func GenDisplaceFn(a,vo,so float64) func (int) float64 {
	distFn := func(t int) float64{
		return so + (vo * float64(t)) + (0.5 * a * math.Pow(float64(t),2.0))
	}
	return distFn
}

func main() {
	var a,vo,so float64
	var t int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter value for acceleration.")
	if !scanner.Scan(){
		return
	}
	a,err := strconv.ParseFloat(scanner.Text(),64)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("enter value for initial velocity.")
	if !scanner.Scan(){
		return
	}
	vo,err = strconv.ParseFloat(scanner.Text(),64)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("enter value for initial displacement.")
	if !scanner.Scan(){
		return
	}
	so,err = strconv.ParseFloat(scanner.Text(),64)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("enter value for time in seconds.")
	if !scanner.Scan(){
		return
	}
	t,err = strconv.Atoi(scanner.Text())
	if err != nil{
		log.Fatal(err)
	}

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println("displacement travelled after time ",t," is ",fn(t))
}
