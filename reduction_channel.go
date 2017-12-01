package main

import (
	"fmt"
	"math/rand"
)

func main() {
	reductionChan := make(chan int)

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

			reductionChan <- max
		}(i)
	}

	globalMax := 0
	for i := 0; i < 10; i++ {
		localMax := <-reductionChan

		if localMax > globalMax {
			globalMax = localMax
		}
	}

	fmt.Println("GLOBAL MAX = ", globalMax)
}