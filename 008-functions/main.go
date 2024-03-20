package main

import (
	"fmt"
)

type transformFn func(int) int

type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	fmt.Println("Hello, World!")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	doubledNumbers := doubleNumber(&numbers)
	fmt.Println(doubledNumbers)

	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tripledNumbers)

	// numbers := []int{1, 2, 3, 4, 5}
	// doubledNumbers := transformNumbers(numbers, double)
	// fmt.Println(doubledNumbers)

	moreNumbers := []int{5,1,2,3}
	transfomerFN1 := getTransformerFunction(&numbers)
	transfomerFN2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transfomerFN1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transfomerFN2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func getTransformerFunction(numbers *[]int) transformFn {
	if(*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}

func doubleNumber(numbers *[]int) []int {
	doubledNumbers := []int{}
	for _, number := range *numbers {
		doubledNumbers = append(doubledNumbers, double(number))
	}

	return doubledNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}