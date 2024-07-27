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
	noise      string
}

func (an *Animal) Eat() {
	fmt.Println(an.food)
}

func (an *Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an *Animal) Speak() {
	fmt.Println(an.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hss"}

	for {
		fmt.Println()
		fmt.Println("Enter animal and an info. For example : 'cow eat'")
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputArr := strings.Split(input, " ")

		name := inputArr[0]
		name = strings.TrimSpace(name)

		info := inputArr[1]
		info = strings.TrimSpace(info)

		if name == "cow" {
			switch info {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		} else if name == "bird" {
			switch info {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		} else if name == "snake" {
			switch info {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}
		} else {
			fmt.Println("Invalid Input")
			fmt.Println("Animals can only be : cow, bird, snake")
			fmt.Println("Info can only be : eat, speak, move")
		}
	}
}
