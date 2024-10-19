package main

import "fmt"

func main() {
	var revenue, expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := float64(ebt) / profit

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("Ratio: %v", ratio)
}