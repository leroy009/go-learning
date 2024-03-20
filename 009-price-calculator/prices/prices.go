package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64 
}

func (job TaxIncludedPriceJob) LoadData() {
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


}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("price %.2f", price)] = price * (1 + job.TaxRate)
	}	

	fmt.Println(result)
}

func NewTaxIncludedPriceJob (taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		InputPrices: []float64{10,20,30,33.5},
		TaxRate: taxRate,
	}
}