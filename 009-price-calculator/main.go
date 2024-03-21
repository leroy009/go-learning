package main

import (
	"leroy.africa/leroy/price-calulator/cmdmanager"
	"leroy.africa/leroy/price-calulator/prices"
)

func main() {
	taxRates := []float64{0.10, 0.15, 0.20, 0.25}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))

		cmdm := cmdmanager.New()

		// priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate) // input from file.
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate) // input from command line.
		priceJob.Process()
	}
}