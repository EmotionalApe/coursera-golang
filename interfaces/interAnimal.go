package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	food       string
	locomotion string
	sound      string
}

func (cw cow) Eat() {
	fmt.Println(cw.food)
}
func (cw cow) Move() {
	fmt.Println(cw.locomotion)
}
func (cw cow) Speak() {
	fmt.Println(cw.sound)
}

type bird struct {
	food       string
	locomotion string
	sound      string
}

func (brd bird) Eat() {
	fmt.Println(brd.food)
}
func (brd bird) Move() {
	fmt.Println(brd.locomotion)
}
func (brd bird) Speak() {
	fmt.Println(brd.sound)
}

type snake struct {
	food       string
	locomotion string
	sound      string
}

func (snk snake) Eat() {
	fmt.Println(snk.food)
}
func (snk snake) Move() {
	fmt.Println(snk.locomotion)
}
func (snk snake) Speak() {
	fmt.Println(snk.sound)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	map1 := make(map[string]Animal)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		inputArr := strings.Fields(input)

		if inputArr[0] == "newanimal" {
			animName := inputArr[1]
			animType := inputArr[2]
			var animal Animal
			switch animType {
			case "cow":
				animal = cow{"grass", "walk", "moo"}

			case "bird":
				animal = bird{"worms", "fly", "peep"}

			case "snake":
				animal = snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Invalid animal type")
				continue
			}

			map1[animName] = animal
			fmt.Println("Created it!")

		} else if inputArr[0] == "query" {
			animName := inputArr[1]
			animInfo := inputArr[2]

			animal, exists := map1[animName]
			if !exists {
				fmt.Println("Animal not found")
				continue
			}
			switch animInfo {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid info type")
			}
		}
	}
}
