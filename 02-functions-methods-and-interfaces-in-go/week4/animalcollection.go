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

type Cow struct {
}

func (aCow *Cow) Eat() {
	fmt.Printf("grass")
}

func (aCow *Cow) Speak() {
	fmt.Printf("moo")
}

func (aCow *Cow) Move() {
	fmt.Printf("walk")
}

type Bird struct {
}

func (aBird *Bird) Eat() {
	fmt.Printf("worms")
}

func (aBird *Bird) Speak() {
	fmt.Printf("peep")
}

func (aBird *Bird) Move() {
	fmt.Printf("fly")
}

type Snake struct {
}

func (aSnake *Snake) Eat() {
	fmt.Printf("mice")
}

func (aSnake *Snake) Speak() {
	fmt.Printf("hsss")
}

func (aSnake *Snake) Move() {
	fmt.Printf("slither")
}

func main() {

	mapOfAnimals := make(map[string]Animal)

	for {
		fmt.Printf("> ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			userInput := strings.ToLower(scanner.Text())
			words := strings.Fields(userInput)
			if len(words) != 3 {
				fmt.Printf("Incorrect number of arguments, was %d, expected 3\n", len(words))
				continue
			}

			switch words[0] {
			case "newanimal":
				myAnimal := CreateAnimal(words[2])
				mapOfAnimals[words[1]] = myAnimal
				fmt.Println("Created it!")
			case "query":
				myAnimal := mapOfAnimals[words[1]]
				if myAnimal == nil {
					fmt.Printf("Could not find an animal named '%s' in the list\n", words[1])
					continue
				} else {
					PrintAction(myAnimal, words[2])
					fmt.Println("")
				}
			default:
				fmt.Printf("Unexpected command '%s'", words[0])
				continue
			}
		} else {
			fmt.Printf("Error occurred: %s", scanner.Err())
		}
	}
}

func CreateAnimal(animalType string) Animal {

	var myAnimal Animal

	switch animalType {
	case "cow":
		myAnimal = new(Cow)
	case "bird":
		myAnimal = new(Bird)
	case "snake":
		myAnimal = new(Snake)
	default:
		fmt.Printf("Unrecognized animal type '%s'", animalType)
	}

	return myAnimal
}

func PrintAction(anAnimal Animal, property string) {

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