package main

import "fmt"

func main () {
	age := 70

	fmt.Println("Age :", age)

	adultYears := getAdultYears(age)
	fmt.Println("Adult Age :", adultYears)
}

func getAdultYears(age int) int {
	return age -18
}