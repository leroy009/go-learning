package main

import (
	"errors"
	"fmt"
	"os"
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
	taxRate, err := getuserInputProper("Input your Tax Rate :")
	
	if err != nil {
		// fmt.Println(err)
		// return
		panic(err)
	}

	// ebt := revenue - expenses
	// profit := ebt * (1- (taxRate/100))
	// ratio := ebt/profit	
	ebt, profit, ratio := calcualteProfits(revenue, expenses, taxRate)

	fmt.Println("Earniong before Tax: ", ebt)
	fmt.Println("Earnings after Tax: ", profit)
	fmt.Println("Ratio: ", ratio)

	writeProfitDataIntoFile(ebt, profit, ratio)

	fmt.Printf("Earnings before Tax: %.2f \nEarnings after Tax: %.2f \nRatio: %.3f", ebt, profit, ratio)
}

func getUserInputErrorProne(message string, value *float64) {
	fmt.Print(message + " :")
	fmt.Scan(value)
}

func getuserInputProper(message string) (float64, error) {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must ne a positive number")
	}

	return userInput, nil
}

func getUserInput(message string, value *float64) {
  fmt.Print(message + " :")
  _, err := fmt.Scan(value)
  if err != nil {
    fmt.Println("Error: Invalid input. Please enter a number.")
    getUserInput(message, value) // Recursively call the function to get valid input
  }
  	if *value <= 0 {
		fmt.Println("Value cannot be '0' or a negative number")
		getUserInput(message, value) // Recursively call the function to get valid input
		return // Ensure we don't continue execution after getting the correct input
	}
}

func calcualteProfits(revenue, expense, taxRate float64) (ebt, profit, ratio float64 ) {
	ebt = revenue - expense
	profit = ebt * (1- (taxRate/100))
	ratio = ebt/profit	

	return ebt, profit, ratio
}

func writeProfitDataIntoFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRation: %.2f", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(data), 0644)
}