package main

import (
"bufio"
"fmt"
"os"
"strings"
)

func main() {

	for {
		fmt.Printf("> ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			userInput := strings.ToLower(scanner.Text())
			words := strings.Fields(userInput)
			if len(words) != 2 {
				fmt.Printf("Incorrect number of arguments, was %d, expected 2\n", len(words))
				continue
			}
			myAnimal := CreateAnimal(words[0])
			PrintAction(&myAnimal, words[1])
		} else {
			fmt.Printf("Error occurred: %s", scanner.Err())
		}
	}
}

type Animal struct {
	food string
	locomotion string
	noise string
}

func (anAnimal *Animal) Eat() {
	fmt.Println(anAnimal.food)
}

func (anAnimal *Animal) Move() {
	fmt.Println(anAnimal.locomotion)
}

func (anAnimal *Animal) Speak() {
	fmt.Println(anAnimal.noise)
}

func CreateAnimal(animalType string) Animal {

	var myAnimal Animal

	switch animalType {
	case "cow":
		myAnimal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		myAnimal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		myAnimal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		fmt.Printf("Unrecognized animal type '%s'", animalType)
	}

	return myAnimal
}

func PrintAction(anAnimal *Animal, property string) {

	switch property {
	case "eat":
		anAnimal.Eat()
	case "move":
		anAnimal.Move()
	case "speak":
		anAnimal.Speak()
	default:
		fmt.Printf("Unrecognized action '%s'", property)
	}
}