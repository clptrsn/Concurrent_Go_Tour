package main

import "fmt"

func create(sender chan string)
{
	sender <- "Hello World"					// Syntax for the "sender" variable to get message to send.
}

func main()
{
	sender := make(chan string)				// Syntax for the "sender" variable to be initialized and set as a string channel.
	
	go create(sender)

	receiver := <-sender					// Syntax for variable to receive message to print.
	fmt.Println(receiver)					// Syntax for the print statement.
}