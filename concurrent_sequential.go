package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(threadId int) {
			time.Sleep(time.Second * time.Duration(threadId))
			fmt.Println(threadId)
		} (i)
	}

	// Go Routines are automatically stopped when main ends
	// Have a user input block main from exiting
	var input string
	fmt.Scan(&input)
}