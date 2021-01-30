package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	for {
		fmt.Println("")
		fmt.Print("> ")
	}
}
