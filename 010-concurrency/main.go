package main

import (
	"fmt"
	"time"
)

func main () {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you?")
	go greet("Nice to meet you too!")
}

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func slowGreet(name string) {
	time.Sleep(5 * time.Second) // simulating a slow task which takes long to complete
	fmt.Println("Hello, ", name)
}