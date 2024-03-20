package main

import "fmt"

func main() {
	var prices []float64 = []float64{10,20,30,33.5}
	taxRates := []float64{0.10, 0.15, 0.20, 0.25}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxInlcudedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxInlcudedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxInlcudedPrices
	}

	fmt.Println(result)

}