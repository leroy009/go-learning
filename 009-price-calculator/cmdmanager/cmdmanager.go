package cmdmanager

import "fmt"

type CMDmanager struct {}

func (cmd CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices, confirm every price with enter")

	var prices []string

	for {
		var price string
		fmt.Println("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)
		
		fmt.Println("Do you want to enter another price? (y/n)")
		var answer string
		fmt.Scan(&answer)
		if answer == "n" {
			break
		}
	}

	return prices, nil
}
func (cmd CMDmanager) WriteResults( data interface {}) error {
	fmt.Println(data)
	return nil

}

func New() CMDmanager {
	return CMDmanager{}
}