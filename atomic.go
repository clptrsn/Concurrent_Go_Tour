package main

import(
	"fmt"
	"sync"
	"sync/atomic"
)

var wg *sync.WaitGroup

func aa(atomAdd *int64, i int64){
	// Adds i to the value inside atomAdd atomically
	atomic.AddInt64(atomAdd, i)
	wg.Done()
}

func main(){
	var atomAdd int64 = 0
	var i int64

	wg = &sync.WaitGroup{}
	
	for i = 0; i < 10; i++ {
		wg.Add(1)
		go aa(&atomAdd, i)
	}
	
	wg.Wait()	// Waits for all goroutines to end
	finalNumber := atomic.LoadInt64(&atomAdd)
	fmt.Println("The Final Number is ", finalNumber)
}