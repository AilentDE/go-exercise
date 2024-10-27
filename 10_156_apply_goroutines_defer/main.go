package main

import (
	"fmt"
	"time"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("start:",time.Now())
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for taxIndex, taxRate := range taxRates {
		doneChans[taxIndex] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process()

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
		go priceJob.Process(doneChans[taxIndex], errorChans[taxIndex])
	}

	// for _, doneChan := range doneChans {
	// 	<- doneChan
	// }
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Printf("[%d] Done!\n", index)
		}
	}
	
	fmt.Println("end:",time.Now())
}
