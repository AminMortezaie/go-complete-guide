package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	hobbies := [3]string{"a", "b", "c"}
	fmt.Printf("%v\n", hobbies)
	fmt.Printf("%v\n", hobbies[0])
	fmt.Printf("%v\n", hobbies[1:3])

	courseGoals := []string{"goal1"}
	courseGoalsClone := append(courseGoals, "goal2")
	fmt.Printf("%v", courseGoalsClone)

	products := []Product{}
	products = append(products, Product{"pro1", 1, 10000})
	fmt.Printf("%v\n", products)

	products = append(products, Product{"pro2", 2, 20000})
	fmt.Printf("%v\n", products)
}
