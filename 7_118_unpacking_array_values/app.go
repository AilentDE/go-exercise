package main

import "fmt"

func main() {
	prices := []float64{10.99, 9.99}
	prices[1] = 8.99

	prices = append(prices, 7.99, 6.99, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{4.99, 3.99, 2.99, 1.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}