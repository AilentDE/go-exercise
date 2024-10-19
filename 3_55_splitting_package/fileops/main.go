package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find this file")
	}
	textString := string(data)
	textFloat, err := strconv.ParseFloat(textString, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored number value")
	}

	return textFloat, nil
}

func WriteFloatToFile(number float64, fileName string) error {
	numberString := fmt.Sprint(number)
	err := os.WriteFile(fileName, []byte(numberString), 0644)
	if err != nil {
		return errors.New("fail to save number into file")
	}
	return nil
}