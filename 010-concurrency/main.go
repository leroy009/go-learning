package main

import (
	"fmt"
	"time"
)

func main () {
	//dones := make([]chan bool, 4)
	done := make(chan bool)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you?", done)
	// dones[3] = make(chan bool)
	go greet("Nice to meet you too!", done)

	// for _, done := range dones {
	// 	<- done
	// }

	for range done {
		
	}
}

func mainUsingChannels1 () {
	dones := make([]chan bool, 4)
	// done := make(chan bool)
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you?", dones[2])
	dones[3] = make(chan bool)
	go greet("Nice to meet you too!", dones[3])

	for _, done := range dones {
		<- done
	}
}

func greet(name string, doneChan chan bool) {
	fmt.Println("Hello, ", name)
	doneChan <- true
}

func slowGreet(name string, doneChan chan bool) {
	time.Sleep(5 * time.Second) // simulating a slow task which takes long to complete
	fmt.Println("Hello, ", name)

	doneChan <- true
	close(doneChan)
}