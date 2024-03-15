package main

import "fmt"

func main () {
	age := 70

	var agePointer *int
	// agePointer := &age
	agePointer = &age

	fmt.Println("Age :", age)
	fmt.Println("Age Pointer Address Value :", agePointer)
	fmt.Println("Age Pointer Value :", *agePointer)
	
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Age :", adultYears)
	fmt.Println("Age :", age)
	fmt.Println("Age Pointer Value :", *agePointer)
}

func getAdultYears(age *int) int {
	return *age -18
}