package main

import (
	"fmt"
)

func main() {
	// var revenue, expenses, taxRate float64
	var revenue, expenses float64

	// fmt.Print("Input your Revenue: ")
	// fmt.Scan(&revenue)

	getUserInput("Input your Revenue", &revenue)
	
	// fmt.Print("Input your Expenses: ")
	// fmt.Scan(&expenses)
	getUserInput("Input your Expenses", &expenses)
	
	// fmt.Print("Input your Tax Rate: ")
	// fmt.Scan(&taxRate)
	// getUserInput("Input your Tax Rate", &taxRate)
	taxRate:= getuserInputProper("Input your Tax Rate :")

	// ebt := revenue - expenses
	// profit := ebt * (1- (taxRate/100))
	// ratio := ebt/profit	
	ebt, profit, ratio := calcualteProfits(revenue, expenses, taxRate)

	fmt.Println("Earniong before Tax: ", ebt)
	fmt.Println("Earnings after Tax: ", profit)
	fmt.Println("Ratio: ", ratio)

	fmt.Printf("Earnings before Tax: %.2f \nEarnings after Tax: %.2f \nRatio: %.3f", ebt, profit, ratio)
}

func getUserInputErrorProne(message string, value *float64) {
	fmt.Print(message + " :")
	fmt.Scan(value)
}

func getuserInputProper(message string) float64 {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)

	return userInput
}

func getUserInput(message string, value *float64) {
  fmt.Print(message + " :")
  _, err := fmt.Scan(value)
  if err != nil {
    fmt.Println("Error: Invalid input. Please enter a number.")
    getUserInput(message, value) // Recursively call the function to get valid input
  }
}

func calcualteProfits(revenue, expense, taxRate float64) (ebt, profit, ratio float64 ) {
	ebt = revenue - expense
	profit = ebt * (1- (taxRate/100))
	ratio = ebt/profit	

	return ebt, profit, ratio
}