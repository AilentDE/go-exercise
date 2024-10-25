package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	result1 := apply(&numbers, double)
	result2 := apply(&numbers, triple)

	fmt.Printf("doubled: %v\ntruple: %v\n", result1, result2)
}

func apply(numberList *[]int, transform func(int) int) []int {
	result := []int{}

	for _, val := range *numberList{
		result = append(result, transform(val))
	}
	return result
}

func createTransformer(factor int) func(int) int {
	fmt.Printf("created function with %d !!\n", factor)

	return func(number int) int {
		return number * factor
	}
}