package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalance() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceString := string(data)
	balance, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalance(balance float64) error {
	balanceString := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceString), 0644)
	if err != nil {
		return errors.New("fail to save balance file")
	}
	return nil
}

func main() {
	accountBalance, err := getBalance()
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}

	// for i := 0; i < 2; i++ {
	for {

		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")

		fmt.Println("---------------------")
		fmt.Println("| 1. Check balance  |")
		fmt.Println("| 2. Deposit money  |")
		fmt.Println("| 3. Withdraw money |")
		fmt.Println("| 4. Exit           |")
		fmt.Println("---------------------")

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
			writeBalance(accountBalance)
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
			writeBalance(accountBalance)
		default:
			fmt.Println("Goodbye.")
			fmt.Printf(`
			Thanks for choosing our bank.
			`)
			return
		}

		// if-else mode
		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated with new amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Your withdraw amount:")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}
			
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You cant't withdraw more than your have.")
		// 		continue
		// 	}
			
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated with new amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye.")
		// 	break
		// }

		fmt.Println()
	}
	// fmt.Printf(`
	// Thanks for choosing our bank.
	// `)
}