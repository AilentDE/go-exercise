package main

import "fmt"

//dynamic array
func main() {
	prices := []float64{10.99, 8.99}
	prices = append(prices, 13.14)
	fmt.Println(prices)
}

//static array
// func main() {
// 	prices := [4]float64{1.1, 2.2, 3.14, 4}
// 	fmt.Println("Array: ", prices)

// 	// slice
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println("Slice: ", highlightedPrices)
// 	fmt.Println("Slice length: ", len(highlightedPrices), "with ori options", cap(highlightedPrices))
// 	highlightedPrices = featuredPrices[1:3]
// 	fmt.Println("Slice length: ", len(highlightedPrices), "with ori options", cap(highlightedPrices))
// }