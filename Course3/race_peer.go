package main

import (
	"fmt"
	"sync"
)

func main() {
	var num int
	var wg sync.WaitGroup
	for i:=0;i<3000;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			num++
		}()
	}
	wg.Wait()
	fmt.Print(num)
}
//what the race condition is: Outcomes of program depends on interleaving.
//Why it occurs: Suppose multiple tasks share a information(like variable). If there isn't any instructions
//				for executing tasks, order of executing tasks become different every time.
//				So, result can't be deterministic.