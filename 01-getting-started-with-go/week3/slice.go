package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Create empty integer slice of length 0 and capacity 3
	intSlice := make([]int, 0, 3)

	var userInput string

	for {
		fmt.Printf("Please enter an integer to add to the slice, or 'X' to exit: ")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Printf("An error occurred processing your input, exiting program...")
			break
		}
		// Did the user ask to exit the program?
		if userInput == "X" {
			fmt.Printf("Here's your sorted slice: %v", intSlice)
			break
		} else {
			// Convert user input to integer, try to add input to the slice, sort the slice
			inputAsInt, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Printf("Couldn't convert user input '%s' to an integer, moving on...\n", userInput)
			} else {
				intSlice = append(intSlice, inputAsInt)
				sort.Ints(intSlice)
			}
		}
	}
}
