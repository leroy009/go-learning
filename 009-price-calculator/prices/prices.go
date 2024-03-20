package prices

import (
	"fmt"

	"leroy.africa/leroy/price-calulator/conversion"
	"leroy.africa/leroy/price-calulator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64 
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")

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

	fmt.Println(result)
}

func NewTaxIncludedPriceJob (taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		InputPrices: []float64{10,20,30,33.5},
		TaxRate: taxRate,
	}
}