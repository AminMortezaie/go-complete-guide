package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/aminmortezaie/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUsersMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)
	return
}

func getUserInput(promptText string) (enteredValue float64) {
	fmt.Print(promptText)
	// underscore means: we don't care about second output.
	userInput, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}
	enteredValue, _ = strconv.ParseFloat(userInput, 64)
	return enteredValue

}
