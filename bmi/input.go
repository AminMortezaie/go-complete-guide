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

	fmt.Print(info.WeightPrompt)
	// underscore means: we don't care about second output.
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		weightInput = strings.Replace(weightInput, "\r\n", "", -1)
		heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	} else {
		weightInput = strings.Replace(weightInput, "\n", "", -1)
		heightInput = strings.Replace(heightInput, "\n", "", -1)
	}

	weight, _ = strconv.ParseFloat(weightInput, 64)
	height, _ = strconv.ParseFloat(heightInput, 64)
	return
}
