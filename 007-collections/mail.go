package main

import "fmt"

type Prodcut struct {
	id string
	name string
	price float64
}

func main() {
	prices := [5]float64{15.59, 12.50, 9.59, 5.99}
	fmt.Println(prices)
}