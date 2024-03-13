package main

import "fmt"

func main () {
	var accountBlace float64 = 1500
	
	fmt.Println("Welcome to Leroy Bank")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check your balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	option := getUserInput("Enter your option: ")


	if option == 1 {
		fmt.Printf("Your account balance is %.2f \n", accountBlace)
	} else if option == 2 {
		amountToDeposit := getUserInput("Enter the amount to deposit: ")

		if amountToDeposit <= 0 {
			fmt.Println("Invalid amount")
			return
		}

		accountBlace = accountBlace + amountToDeposit
		fmt.Printf("Your new account balance is %.2f \n", accountBlace)
	} else if option == 3 {
		amountToWithdraw := getUserInput("Enter the amount to withdraw: ")

		if amountToWithdraw <= 0 {
			fmt.Println("Invalid amount")
			return
		}
		
		if amountToWithdraw > accountBlace {
			fmt.Printf("Invalid amount. You can't withdraw more your account balance: %.2f\n", accountBlace)
			return
		}

		accountBlace = accountBlace - amountToWithdraw
		fmt.Printf("Your new account balance is %.2f \n", accountBlace)
	} else if option == 4 {
		fmt.Println("Goodbye!")	
	}



	fmt.Printf("Your option: %v", option)
}

func getUserInput(message string) float64 {
	var userInput float64
	fmt.Print(message)

	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a number.")
		getUserInput(message) // Recursively call the function to get valid input
	}

	return userInput
}

