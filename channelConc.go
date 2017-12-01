package main

import (
	"fmt"
	"time"
	"math/rand"
)

func work1(done1 chan bool) {

	time.Sleep(time.Second * time.Duration(rand.Intn(5) + 1))
	
	fmt.Println("Worker 1 starting.")
	time.Sleep(time.Second * time.Duration(rand.Intn(5) + 1))
	
	fmt.Println("Worker 1 is done.")
	time.Sleep(time.Second)
	
	fmt.Println("Worker 1 is turning in work.");
	
	
	done1 <- true																	// Sending message that worker 1 is completed.
}

func work2(done2 chan bool) {

	time.Sleep(time.Second * time.Duration(rand.Intn(5) + 1))
	
	fmt.Println("Worker two starting.")
	time.Sleep(time.Second * time.Duration(rand.Intn(5) + 1))
	
	fmt.Println("Worker two is halfway done.")
	time.Sleep(time.Second * time.Duration(rand.Intn(5) + 1))
	
	fmt.Println("Worker two is now done.")
	time.Sleep(time.Second)
	
	fmt.Println("Worker two is now turning in work.")
	
	done2 <- true																	// Sending message that worker 2 is completed.
}

func main() {

	done1 := make(chan bool)
	done2 := make(chan bool)
	
	go work1(done1)																	//Start worker 1.
	go work2(done2)																	//Start worker 2.
	
	// Lets assume worker 2's work needs to be done first
	<-done2																			// Collect "work from worker 2.
	fmt.Println("Worker 2's work has been received, and is now being analyzed.")
	time.Sleep(time.Second*2)
	
	fmt.Println("Worker 2's work has been analyzed. Collecting from Worker 1.")		// Collect "work" from worker 1 after worker 2.
	<-done1
	fmt.Println("Worker 1's work has been received, and is now being analyzed.")
	time.Sleep(time.Second)
	
	fmt.Println("Both workers' work have been received and analyzed. Terminating program.")
}

// Added multiple sleep parameters to help with possible randomness between work loads, ultimately needing worker 2 to be completed first at the end of the program before accepting "work" from worker 1.