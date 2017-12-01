package main

import (
	"fmt"
	"time"
)

func employee(finish chan bool) {	
	fmt.Print("Doing work...")
	time.Sleep(time.Second*2)
	
	fmt.Print("doing more work...")
	time.Sleep(time.Second*2)
	
	fmt.Println("completed work for the day.")
	time.Sleep(time.Second)
	
	fmt.Println("Taking nap before notifying that work is complete.")
	time.Sleep(time.Second*5)
	
	finish <- true			// Giving "fin" the message that it the worker is complete. Added multiple sleep parameters to show that the program will not terminate without this message.
}

func main() {
	finish := make(chan bool, 1)
	go employee(finish)

	<-finish				// Until "fin" has a message to be received, telling us that the worker is done, the program will not terminate.
	
	fmt.Println("Employee finally done with work.")
}