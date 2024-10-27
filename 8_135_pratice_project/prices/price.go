package prices

import (
	"fmt"
	"pratice-project/conversion"
	"pratice-project/iomanager"
)

type TaxIncludedPriceJob struct {
	IOmanager iomanager.IOManager `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (t *TaxIncludedPriceJob) LoadData() error {
	lines, err := t.IOmanager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}
	t.InputPrices = prices
	return nil
}

func (t *TaxIncludedPriceJob) Process() error {
	err := t.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range t.InputPrices {
		caPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("[%.2f]", price)] = fmt.Sprintf("%6.2f", caPrice)
	}

	t.TaxIncludedPrices = result
	return t.IOmanager.WriteJSON(t)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, inputTaxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOmanager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     inputTaxRate,
	}
}