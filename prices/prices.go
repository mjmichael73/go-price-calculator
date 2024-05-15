package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprint(price)] = price * (1 + job.TaxRate)
	}
	result[taxRate] = taxIncludedPrices
}
