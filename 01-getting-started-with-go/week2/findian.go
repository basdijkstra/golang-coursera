package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input a string: ")
	if scanner.Scan() {
		userInputLowerCase := strings.ToLower(scanner.Text())
		if strings.HasPrefix(userInputLowerCase, "i") && strings.Contains(userInputLowerCase, "a") && strings.HasSuffix(userInputLowerCase, "n") {
			fmt.Printf("Found!")
		} else {
			fmt.Printf("Not Found!")
		}
	} else {
		fmt.Printf("Error occurred: %s", scanner.Err())
	}
}
