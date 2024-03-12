package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6.5

func main() {
	var investmentAmount, years float64 = 1200, 5
	var expectedReturnRate float64 

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)


	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)


	// futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	formartedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formartedRFV := fmt.Sprintf("Future Value (adjusted based on inflation): %.2f\n", futureRealValue)
	fmt.Print(formartedFV, formartedRFV)

	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted based on inflation): %.2f", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.2f
// Future Value (adjusted based on inflation): %.2f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, realFutireValue float64) {
	// futureValue:= investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years) 
	// realFutireValue := futureValue / math.Pow(1 + inflationRate/100, years)
	// return futureValue, realFutireValue
	futureValue = investmentAmount * math.Pow((1 + expectedReturnRate/ 100), years) 
	realFutireValue = futureValue / math.Pow(1 + inflationRate/100, years)	
}