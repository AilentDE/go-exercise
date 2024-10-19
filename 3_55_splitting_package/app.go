package main

import (
	"fmt"

	"example.com/splitting-package/fileops"
)

const accountBalanceFile = "balance.txt"



func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}

	for {
		fmt.Println("Welcome to Go Bank!")
		presentOptions()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		// switch mode
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated with new amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdraw amount:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cant't withdraw more than your have.")
				continue
			}
			
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated with new amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye.")
			fmt.Printf(`
			Thanks for choosing our bank.
			`)
			return
		}

		fmt.Println()
	}
}