package main

import "fmt"

func main() {

	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	prices[1] = 9.99
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

}

// func main() {

// 	var productNames [4]string = [4]string{"A book"}

// 	fmt.Println(productNames)

// 	prices := [4]float64{10.99, 20.0, 45.99, 88.99}
// 	prices[2] = 66
// 	fmt.Println(prices)
// 	featuredPrices := prices[:3]
// 	fmt.Println(featuredPrices)

// 	// there is no way to trace the left side of an array and if you initiate your slice from second object of your array you cannot achieve to your first object neither.
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// }
