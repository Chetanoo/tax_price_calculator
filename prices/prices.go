package prices

import (
	"fmt"
	"price_calculator/conversion"
	"price_calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IOManager:   iom,
	}
}

type ResultMap map[string]string

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	resultMap := make(ResultMap)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		resultMap[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = resultMap
	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}
