package main

import "fmt"

func main() {
	revenue := inputFloat64("Revenue: ")
	expenses := inputFloat64("Expenses: ")
	taxRate := inputFloat64("Tax Rate: ")

	concluteResult(revenue, expenses, taxRate)
}

func inputFloat64(text string) float64 {
	fmt.Print(text)
	var temp float64
	fmt.Scan(&temp)
	return temp
}

func concluteResult(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("===== Result =====")
	fmt.Printf("%v\n%v\n%.2f", ebt, profit, ratio)

	return ebt, profit, ratio
}