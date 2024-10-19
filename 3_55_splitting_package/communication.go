package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOptions() {

	fmt.Println("What do you want to do?")

	fmt.Println("---------------------")
	fmt.Println("| 1. Check balance  |")
	fmt.Println("| 2. Deposit money  |")
	fmt.Println("| 3. Withdraw money |")
	fmt.Println("| 4. Exit           |")
	fmt.Println("---------------------")
	fmt.Println("or reach us 24/7", randomdata.PhoneNumber())
}