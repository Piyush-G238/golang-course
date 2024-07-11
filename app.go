package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// sending data to channel
	doneChan <- true
}

func main() {

	isdone := make(chan bool)

	// go greet("Nice to meet you!")
	// go greet("How are you?")
	go slowGreet("How ... are ... you ...?", isdone)

	go greet("I hope you're liking the course!")
	<-isdone
}
