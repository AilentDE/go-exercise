package cmdmanager

import (
	"errors"
	"fmt"
	"strconv"
)

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER.")

	var prices []string
	
	isContinue := true
	for isContinue {
		var input string
		fmt.Print("Price: ")
		fmt.Scan(&input)
		inputVal, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return nil, errors.New("something while enter value")
		}

		if  inputVal <= 0 {
			isContinue = false
		} else {
			prices = append(prices, input)
		}
	}

	return prices, nil
}

func (cmd CMDManager) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}