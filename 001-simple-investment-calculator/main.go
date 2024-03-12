package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1200
	var expectedReturnRate float64 = 12.36
	var years int = 5

	var futureValue = investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years)

	fmt.Println(futureValue)
}