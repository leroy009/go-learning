package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Input your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input your Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1- (taxRate/100))
	ratio := ebt/profit
	

	fmt.Println("Earniong before Tax: ", ebt)
	fmt.Println("Earnings after Tax: ", profit)
	fmt.Println("Ratio: ", ratio)

	fmt.Printf("Earnings before Tax: %v \nEarnings after Tax: %v \nRatio: %v", ebt, profit, ratio)

}
