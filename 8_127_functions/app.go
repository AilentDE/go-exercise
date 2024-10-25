package main

import "fmt"

type changeNumberFnc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transform(&numbers, double)
	fmt.Println("double to:", doubled)

	tripled := transform(&numbers, triple)
	fmt.Println("Triple to:", tripled)

	// more nest
	moreNumbers1 := []int{1, 2, 3, 4, 5}
	moreNumbers2 := []int{6, 7, 8, 9, 10}

	func1 := getTransformFuncion(&moreNumbers1)
	func2 := getTransformFuncion(&moreNumbers2)

	result1 := transform(&moreNumbers1, func1)
	result2 := transform(&moreNumbers2, func2)

	fmt.Printf("func1 result: %v\nfunc2 result: %v\n", result1, result2)
}

func transform(numberList *[]int, apply changeNumberFnc) []int {
	transformed := []int{}

	for _, val := range *numberList {
		transformed = append(transformed, apply(val))
	}

	return transformed
}

func getTransformFuncion(sampleList *[]int) changeNumberFnc {
	if (*sampleList)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}