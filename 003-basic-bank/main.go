package main

import (
	"fmt"

	"leroy.africa/leroy/simple-bank/fileops"
)

const accountFile = "data.txt"
func main () {
	accountBalance, err := fileops.GetDataFromFile(accountFile)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		fmt.Println("--------------------------------------")
	}

	fmt.Println("Welcome to Leroy Bank")

	for {		
		presentOption()
		option := getUserInput("Enter your option: ")
		
		switch option {
			case 1:
				fmt.Printf("Your account balance is %.2f \n", accountBalance)
			case 2:
				amountToDeposit := getUserInput("Enter the amount to deposit: ")

				if amountToDeposit <= 0 {
					fmt.Println("Invalid amount")
					continue
				}

				accountBalance = accountBalance + amountToDeposit
				fmt.Printf("Your new account balance is %.2f \n", accountBalance)
				fileops.WriteDataToFile(accountFile, accountBalance)
			case 3:
				amountToWithdraw := getUserInput("Enter the amount to withdraw: ")

				if amountToWithdraw <= 0 {
					fmt.Println("Invalid amount")
					continue
				}

				if amountToWithdraw > accountBalance {
					fmt.Printf("Invalid amount. You can't withdraw more your account balance: %.2f\n", accountBalance)
					continue
				}

				accountBalance = accountBalance - amountToWithdraw
				fmt.Printf("Your new account balance is %.2f \n", accountBalance)
				fileops.WriteDataToFile(accountFile, accountBalance)
			case 4:
				fmt.Println("Goodbye!")	
				return
				fmt.Println("Thank you for using this good bank")
			default:
				continue
		}
		/*
		if option == 1 {
			fmt.Printf("Your account balance is %.2f \n", accountBlace)
		} else if option == 2 {
			amountToDeposit := getUserInput("Enter the amount to deposit: ")

			if amountToDeposit <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBlace = accountBlace + amountToDeposit
			fmt.Printf("Your new account balance is %.2f \n", accountBlace)
		} else if option == 3 {
			amountToWithdraw := getUserInput("Enter the amount to withdraw: ")

			if amountToWithdraw <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if amountToWithdraw > accountBlace {
				fmt.Printf("Invalid amount. You can't withdraw more your account balance: %.2f\n", accountBlace)
				continue
			}

			accountBlace = accountBlace - amountToWithdraw
			fmt.Printf("Your new account balance is %.2f \n", accountBlace)
		} else if option == 4 {
			fmt.Println("Goodbye!")	
			break
		}
		fmt.Printf("Your option: %v", option)
		*/
	} 

	fmt.Println("Thank you for using this good bank")
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



