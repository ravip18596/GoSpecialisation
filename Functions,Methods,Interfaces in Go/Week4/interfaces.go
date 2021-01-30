package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cow struct {
	name string
}

func (c cow) Eat() {
	fmt.Println(fmt.Sprintf("%s eats grass",c.name))
}

func (c cow) Move() {
	fmt.Println(fmt.Sprintf("%s walk",c.name))
}

func (c cow) Speak() {
	fmt.Println(fmt.Sprintf("%s speaks moo",c.name))
}

type bird struct {
	name string
}

func (b bird) Eat() {
	fmt.Println(fmt.Sprintf("%s eats worms",b.name))
}

func (b bird) Move() {
	fmt.Println(fmt.Sprintf("%s fly",b.name))
}

func (b bird) Speak() {
	fmt.Println(fmt.Sprintf("%s speaks peep",b.name))
}

type snake struct {
	name string
}

func (s snake) Eat() {
	fmt.Println(fmt.Sprintf("%s eats mice",s.name))
}

func (s snake) Move() {
	fmt.Println(fmt.Sprintf("%s moves like slither",s.name))
}

func (s snake) Speak() {
	fmt.Println(fmt.Sprintf("%s speaks hsss",s.name))
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	objMap := make(map[string]Animal)
	for {
		fmt.Println()
		fmt.Println(fmt.Sprintf(`select between "%s" or "%s" command. Each line should contain 3 words starting with the command`, "newanimal", "query"))
		fmt.Println("allowed animal types are ")
		fmt.Println("cow")
		fmt.Println("bird")
		fmt.Println("snake")
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()
		commands := strings.Split(command, " ")
		if len(commands) != 3 {
			fmt.Println("Incorrect command. Please enter 3 string command")
			break
		}
		switch strings.ToLower(commands[0]) {
		case "newanimal":
			var an Animal
			animalName := commands[1]
			objectType := commands[2]
			switch strings.ToLower(objectType) {
			case "cow":
				an = cow{name: animalName}
			case "bird":
				an = bird{name: animalName}
			case "snake":
				an = snake{name: animalName}
			default:
				fmt.Println("error. incorrect object type. allowed types are cow, bird, snake.")
			}
			objMap[animalName] = an
			fmt.Println("Created it!")
		case "query":
			animalName := commands[1]
			action := commands[2]
			var animal Animal
			animal, ok := objMap[animalName]
			if !ok {
				fmt.Println(fmt.Sprintf("animal with name %s not present", animalName))
				break
			}
			switch strings.ToLower(action) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("error. incorrect object type. allowed types are cow, bird, snake.")
			}

		default:
			fmt.Println("error - incorrect first command. select b/w newanimal or query.")
		}
	}
}
