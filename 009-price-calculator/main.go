package main

import (
	"fmt"

	"leroy.africa/leroy/price-calulator/filemanager"
	"leroy.africa/leroy/price-calulator/prices"
)

func main() {
	taxRates := []float64{0.10, 0.15, 0.20, 0.25}
	doneChans := make([]chan bool, len(taxRates))
	erroChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)
		erroChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))

		//cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate) // input from file.
		//priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate) // input from command line.
		go priceJob.Process(doneChans[index], erroChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process jon")
		// 	fmt.Println(err)
		// }
	}
	for _, errorChan := range erroChans {
		<- errorChan
	}
	
	for _, doneChan := range doneChans {
		<- doneChan
	}
}