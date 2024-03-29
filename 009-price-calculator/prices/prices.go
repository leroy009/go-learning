package prices

import (
	"fmt"

	iomanager "leroy.africa/leroy/price-calulator/IOManager"
	"leroy.africa/leroy/price-calulator/conversion"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"` // filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines() //filemanager.ReadLines("prices.txt")

	if err != nil {
		// fmt.Printf("Error: %s\n", err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		// fmt.Println("Converting price to float failed")
		// fmt.Printf("Error: %s\n", err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	// err = errors.New("Could not load data")

	if err != nil {
		// return err
		errorChan <- err
		return
	}

	result := make(map[string]string)
	
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("price %.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}	

	job.TaxIncludedPrices = result
	// errorChan <- job.IOManager.WriteResults(job)
	job.IOManager.WriteResults(job)

	doneChan <- true
}

func NewTaxIncludedPriceJob (iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		IOManager: iom,
		InputPrices: []float64{10,20,30,33.5},
		TaxRate: taxRate,
	}
}