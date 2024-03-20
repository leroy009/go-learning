package main

import "fmt"

type Prodcut struct {
	id string
	name string
	price float64
}

func main () {
	websites := map[string]string{
		"google": "https://www.google.com",
		"facebook": "https://www.facebook.com",
		"twitter": "https://www.twitter.com",
		"amazon web services": "https://www.aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["amazon web services"])

	websites["linkedin"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}





func mainArraysSlicesAndAll() {
	mainDynamicSlices()
	mainExcercise()
	slicesAndArrays()
}

func mainExcercise() {
	// 1) Create a new array (!) that contains three hobbies you have
	hobboes := []string{"Swimming", "Hunting", "Reading"}
	// 		Output (print) that array in the command line.
	fmt.Println(hobboes)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobboes[0])
	//		- The second and third element combined as a new list
	secondAndThird := hobboes[1:3]
	fmt.Println(secondAndThird)
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	firstTwo := hobboes[:2]
	fmt.Println(firstTwo)
	//		Create that slice in two different ways (i.e. create two slices in the end)
	firstTwoAgain := hobboes[0:2]
	fmt.Println(firstTwoAgain)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	secondAndLast := hobboes[1:3]
	fmt.Println(secondAndLast)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learn Go", "Build projects"}
	fmt.Println(goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learn more Go"
	goals = append(goals, "Get a job")
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	products := []Prodcut{
		{id: "p1", name: "Product 1", price: 15.59},
		{id: "p2", name: "Product 2", price: 12.50},
	}
	fmt.Println(products)
	
	//		Then add a th ird product to the existing list of products.
	newProduct := Prodcut{id: "p3", name: "Product 3", price: 9.59}
	products = append(products, newProduct)
	fmt.Println(products)
}

func mainDynamicSlices() {
	prices := []float64{15.59, 12.50, 9.59, 5.99}
	fmt.Println(prices[0])
	fmt.Println(prices[1:2])

	prices = append(prices, 5.99)
	fmt.Println(prices)

	prices = prices[1:]
	fmt.Println(prices)

	discountedPrices := []float64{10.59, 11.50, 8.59, 4.99}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)

}

func slicesAndArrays() {
	var prodNames [4]string = [4]string{"Laptop", "Mouse", "Keyboard"}
	prices := [5]float64{15.59, 12.50, 9.59, 5.99}
	fmt.Println(prices)
	fmt.Println(prodNames)

	fmt.Println(prices[2])

	prodNames[3] = "Monitor"
	fmt.Println(prodNames[3])

	featuredPrices := prices[1:3]
	featuredPrices[0] = 15.60
	fmt.Println(featuredPrices)
	fmt.Println(prices)
	otherfeaturedPrices := prices[:3]
	fmt.Println(otherfeaturedPrices)
	othernewfeaturedPrices := prices[2:]
	fmt.Println(othernewfeaturedPrices)

	highlightedPrices := featuredPrices[:2]
	fmt.Println(highlightedPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}