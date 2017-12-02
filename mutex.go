package main

import(
	"fmt"
	"sync"
	"time"
)

var mutex *sync.Mutex = &sync.Mutex{}
var wg *sync.WaitGroup 
	
func criticalHi(i int){
	// Creates a critical section to wait and print hello
	mutex.Lock()
	time.Sleep(time.Second)
	fmt.Println("Hello from routine ", i)
	mutex.Unlock()

	wg.Done()
}

func main(){
	wg = &sync.WaitGroup{}
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go criticalHi(i)
	}

	wg.Wait() // Wait for all go routines to run
}