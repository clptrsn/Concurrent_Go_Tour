package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var globalMax int = 0
	var mutexLock *sync.Mutex = &sync.Mutex{}

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(threadNum int) {
			// Generate a list of 10 numbers
			var list [10]int
			for n := 0; n < 10; n++ {
				list[n] = rand.Intn(1000)
			}

			fmt.Println(threadNum, ": ", list)

			max := 0
			for n := 0; n < 10; n++ {
				if list[n] > max {
					max = list[n]
				}
			}

			fmt.Println(threadNum, " max ", max)

			mutexLock.Lock()
			if max > globalMax {
				globalMax = max
			}
			mutexLock.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("GLOBAL MAX = ", globalMax)
}