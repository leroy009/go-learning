package main

import "fmt"

func main() {
	fact := factorial(3)
	fmt.Println(fact)

}

func factorial(number int) int {
	// result := 1 
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	// Using recursion
	if number == 0 {
		return 1
	}
	return number * factorial(number - 1)
}