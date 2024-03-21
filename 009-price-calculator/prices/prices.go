package prices

import (
	"fmt"

	"leroy.africa/leroy/price-calulator/conversion"
	"leroy.africa/leroy/price-calulator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines() //filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("Converting price to float failed")
		fmt.Printf("Error: %s\n", err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("price %.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}	

	job.TaxIncludedPrices = result

	job.IOManager.WriteResults(job)
}

func NewTaxIncludedPriceJob (fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		IOManager: fm,
		InputPrices: []float64{10,20,30,33.5},
		TaxRate: taxRate,
	}
}