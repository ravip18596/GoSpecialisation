package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise	   string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func initAnimalObjects() (map[string]*Animal){
	cow := &Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := &Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := &Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	m := make(map[string]*Animal)
	m["cow"] = cow
	m["bird"]= bird
	m["snake"] = snake
	return m
}

func main() {
	objectMap := initAnimalObjects()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("select a animal from cow,bird,snake.")
		fmt.Print("> ")
		scanner.Scan()
		object := strings.ToLower(scanner.Text())
		if _,ok:=objectMap[object];!ok{
			fmt.Println("animal not selected from above list")
			break
		}
		fmt.Println("select option to know about animal - eat,move,speak.")
		fmt.Print("> ")
		scanner.Scan()
		action := strings.ToLower(scanner.Text())
		switch action {
		case "eat":
			objectMap[object].Eat()
		case "move":
			objectMap[object].Move()
		case "speak":
			objectMap[object].Speak()
		default:
			fmt.Println("action not selected from eat,move,speak.")
		}
	}
}
