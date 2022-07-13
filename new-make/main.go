package main

import "fmt"

// using iota to make a list of constants and assign value to them.

// const (
// 	inputAttack = iota
// 	inputSpecialAttack
// 	inputHeal
// )

func main() {

	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	// only for map, slice and channels.
	hobbies := make([]string, 0, 10)
	fmt.Printf("hobbies: %v", cap(hobbies))
}
