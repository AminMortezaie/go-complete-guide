package conditions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	userAge, err := getUserAge()

	if err != nil {
		fmt.Printf("Invalid input!")
		return
	}

	// storing a condition in a variable could be super useful and Go will approve this.
	isOldEnough := userAge >= 18
	isOld := userAge >= 30 && userAge < 50

	if isOld {
		fmt.Println("You're eligible for our senior jobs!")
	} else if isOldEnough {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry, You're not old enough :/")
	}
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter your age: ")

	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.ReplaceAll(userAgeInput, "\r\n", "")
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	// Creating New Error
	// error := errors.New("INVALID INPUT!")
	// return int(userAge), error

	return int(userAge), err
}
