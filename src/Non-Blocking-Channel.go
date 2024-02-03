package main

import "fmt"

func main() {
	messages := make(chan string)
	singlas := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-singlas:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
