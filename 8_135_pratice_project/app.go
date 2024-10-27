package main

import (
	"fmt"
	"pratice-project/cmdmanager"
	"pratice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job.")
			fmt.Println(err)
		}
	}
}