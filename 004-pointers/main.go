package main

import "fmt"

func main () {
	age := 70

	// agePointer := &age
	var agePointer *int
	agePointer = &age

	fmt.Println("Age :", age)
	fmt.Println("Age Pointer Address Value :", agePointer)
	fmt.Println("Age Pointer Value :", *agePointer)
	
	editAgeToAdultYears(agePointer)
	// fmt.Println("Adult Age :", adultYears)
	fmt.Println("Age :", age)
	fmt.Println("Age Pointer Value :", *agePointer)
}

func editAgeToAdultYears(age *int) {
	*age = *age -18
}