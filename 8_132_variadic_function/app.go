package main

import (
	"fmt"
)

func main() {

	fmt.Println(sumup(1, 2, 3, 4, 100, -10))

	anotherSlice := []int{2, 3, 4, 100, -10}
	fmt.Println(sumup(1, anotherSlice...))
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}