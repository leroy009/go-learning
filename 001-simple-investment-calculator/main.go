package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investmentAmount, years float64 = 1200, 5
	expectedReturnRate := 12.36

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}