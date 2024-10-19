package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	// Save results to JSON file
	results := map[string]float64{
		"EBT":    ebt,
		"Profit": profit,
		"Ratio":  ratio,
	}

	if err := saveResultsToJSON("financials.json", results); err != nil {
		fmt.Println("Error saving results to JSON:", err)
	}
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("you must input not negative nor zero number")
	}
	return userInput, nil
}

func saveResultsToJSON(filename string, results map[string]float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(results); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}
	return nil
}