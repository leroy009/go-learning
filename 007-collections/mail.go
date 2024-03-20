package main

import "fmt"

type Prodcut struct {
	id string
	name string
	price float64
}

func main() {
	var prodNames [4]string = [4]string{"Laptop", "Mouse", "Keyboard"}
	prices := [5]float64{15.59, 12.50, 9.59, 5.99}
	fmt.Println(prices)
	fmt.Println(prodNames)

	fmt.Println(prices[2])

	prodNames[3] = "Monitor"
	fmt.Println(prodNames[3])


}