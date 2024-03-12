package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 1200.50, 5.0, 12.36

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years)

	fmt.Println(futureValue)
}