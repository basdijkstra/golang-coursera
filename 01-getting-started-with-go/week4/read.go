package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var fileName string

	personSlice := make([]Person, 0, 3)

	// Prompt for file name
	fmt.Printf("Please enter a file name: ")
	scannerPrompt := bufio.NewScanner(os.Stdin)
	if scannerPrompt.Scan() {
		fileName = scannerPrompt.Text()
	} else {
		fmt.Printf("Error occurred when reading file name from prompt")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error occurred when opening file")
		os.Exit(1)
	}
	// Moved this to a separate function to be able to do error handling properly
	defer Close(file)

	scannerFile := bufio.NewScanner(file)

	// Read each line into a Person struct and append it to the slice
	for scannerFile.Scan() {
		splitLine := strings.Split(scannerFile.Text(), " ")
		personSlice = append(personSlice, Person{splitLine[0],splitLine[1]})
	}
	if err := scannerFile.Err(); err != nil {
		fmt.Printf("Error when reading text from file:\n%s", scannerFile.Err())
	}

	// Print all Persons in the slice
	for _, person := range personSlice {
		fmt.Printf("First name: %s, last name: %s\n", person.fname, person.lname)
	}
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
