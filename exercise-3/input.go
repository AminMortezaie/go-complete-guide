package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserData(promptText string) (cleanedInput string) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	return
}
