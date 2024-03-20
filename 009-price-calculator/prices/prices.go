package prices

import (
	"bufio"
	"fmt"
	"os"

	"leroy.africa/leroy/price-calulator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64 
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err :=os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Printf("Error: %s\n", err)
		return
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file contents failed!")
		fmt.Printf("Error: %s\n", err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("Converting price to float failed")
		fmt.Printf("Error: %s\n", err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
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