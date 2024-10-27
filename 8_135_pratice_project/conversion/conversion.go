package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	convertedFloat := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		convertedFloat[stringIndex] = floatPrice
	}
	return convertedFloat, nil
}