package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investmentAmount, years float64 = 1200, 5
	var expectedReturnRate float64 

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)


	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)


	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted based on inflation): %.2f", futureValue, futureRealValue)
}