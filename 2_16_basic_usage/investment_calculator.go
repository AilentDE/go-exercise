package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount, years float64 = 1000, 10
	// expectedReturnRate := 5.5
	// or
	// var investmentAmount, years, expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	// fmt.Scan(&investmentAmount)
	investmentAmount := outputText("Investment Amount: ")

	// fmt.Print("Expected Return Rate: ")
	// fmt.Scan(&expectedReturnRate)
	expectedReturnRate := outputText("Expected Return Rate: ")

	// fmt.Print("Investment Years: ")
	// fmt.Scan(&years)
	years := outputText("Investment Years: ")

	// futureValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years) 
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := caclulateFeatureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println("====================")

	// simple print
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	// formatted print
	// fmt.Printf("[v] \nFuture Value: %v \nFuture Real Value: %v \n", futureValue, futureRealValue)
	// fmt.Printf("[f] \nFuture Value: %.2f \nFuture Real Value: %.2f \n", futureValue, futureRealValue)

	// formatted string
	// formattedFV := fmt.Sprintf("Future Value: %.2f \n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Real Value: %.2f \n", futureRealValue)
	// fmt.Print(formattedFV, formattedFRV)

	// multiline string
	fmt.Printf(`
	Future Value: %.2f
	Future Real Value: %.2f
	`, futureValue, futureRealValue)
}

func outputText(text string) float64 {
	fmt.Print(text)
	var tmp float64
	fmt.Scan(&tmp)
	return tmp
}

func caclulateFeatureValues(investmentAmount , expectedReturnRate , years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years) 
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}