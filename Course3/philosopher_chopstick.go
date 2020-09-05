package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func delegate(request chan int, permission chan bool, finished chan int) {
	chopsticks := [5]bool{true, true, true, true, true}
	for  {
		select {
		case phil := <-request:
			if chopsticks[phil-1] && chopsticks[phil%5] {
				permission <- true
				chopsticks[phil-1] = false
				chopsticks[phil%5] = false
			} else {
				permission <- false
			}
		case phil := <-finished:
			chopsticks[phil-1] = true
			chopsticks[phil%5] = true
		}

	}
}

func eat(phil int, request chan int, permission chan bool, finished chan int) {
	defer wg.Done()
	for i := 0; i < 3; {
		request <- phil
		if <-permission {
			fmt.Println("starting to eat ",phil)
			fmt.Println("finishing eating ",phil)
			i++
			finished <- phil
		}
	}
}

func main() {
	request := make(chan int)
	permission := make(chan bool)
	finished := make(chan int)
	go delegate(request, permission, finished)
	wg.Add(5)
	for i := 1; i < 6; i++ {
		go eat(i, request, permission, finished)
	}
	wg.Wait()
}
