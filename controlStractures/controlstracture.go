package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter your age: ")

	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.ReplaceAll(userAgeInput, "\r\n", "")
	userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

	// storing a condition in a variable could be super useful and Go will approve this.
	isOldEnough := userAge >= 18
	isOld := userAge >= 30

	if isOld {
		fmt.Println("You're eligible for our senior jobs!")
	} else if isOldEnough {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry, You're not old enough :/")
	}
}
