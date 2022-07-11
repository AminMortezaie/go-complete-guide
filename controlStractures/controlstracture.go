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

	fmt.Println("Please enter your age:")

	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.ReplaceAll(userAgeInput, "\r\n", "")
	userAge, _ = strconv.ParseInt(userAgeInput, 0, 64)

}
