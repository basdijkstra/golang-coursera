package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	personMap := make(map[string]string)

	personMap = scanInput("name", personMap)
	personMap = scanInput("address", personMap)

	byteArr, err := json.MarshalIndent(personMap,"","\t")
	if err != nil {
		fmt.Printf("Error occurred when marshalling map to JSON")
	} else {
		fmt.Println("Here's your JSON:")
		fmt.Println(string(byteArr))
	}
}

func scanInput(fieldName string, personMap map[string]string) map[string]string {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter a value for '%s': ", fieldName)
	if scanner.Scan() {
		userInput := scanner.Text()
		personMap[fieldName] = userInput
	} else {
		fmt.Printf("Error occurred processing user input: %s", scanner.Err())
		os.Exit(1)
	}
	return personMap
}
