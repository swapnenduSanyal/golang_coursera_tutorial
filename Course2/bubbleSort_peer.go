package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"reflect"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a sequence of integers separated by spaces: ")
	scanner.Scan()
	line := scanner.Text()
	numberStrings := strings.Split(line, " ")
	numbers := []int{}
	for _, item := range numberStrings {
		number, parseError := strconv.Atoi(item)
		if (parseError == nil) {
			numbers = append(numbers, number)
		}
	}
	fmt.Printf("Your numbers: \t\t%v\n", numbers)
	BubbleSort(numbers)
	fmt.Printf("Your sorted numbers: \t%v\n", numbers)
}

func BubbleSort(slice []int) {
	var isSorted bool = true
	for {
		for i, item := range slice {
			if (i+1 < len(slice) && item > slice[i+1]) {
				isSorted = false
				Swap(slice, i, i+1)
			}
		}
		if isSorted {
			return 
		}
		isSorted = true
	}
}

func Swap(slice []int, i int, j int) {
	reflect.Swapper(slice)(i, j)
}