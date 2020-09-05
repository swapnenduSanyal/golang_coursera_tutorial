package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func swap(numbers []int, i int) {
	if numbers[i] > numbers[i+1] {
		temp := numbers[i]
		numbers[i] = numbers[i+1]
		numbers[i+1] = temp
	}
}

func bubbleSort(sli []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(sli)-1; i++ {
		swap(sli, i)
	}
	if len(sli) > 1 {
		wg.Add(1)
		bubbleSort(sli[0:len(sli)-1], wg)
	}
}

func merge(sli1, sli2 []int) []int {
	result := make([]int, 0)
	i1 := 0
	i2 := 0
	for i1 < len(sli1) && i2 < len(sli2) {
		if sli1[i1] < sli2[i2] {
			result = append(result, sli1[i1])
			i1++
		} else {
			result = append(result, sli2[i2])
			i2++
		}
	}
	for ; i1 < len(sli1); i1++ {
		result = append(result, sli1[i1])
	}
	for ; i2 < len(sli2); i2++ {
		result = append(result, sli2[i2])
	}
	return result
}

func main() {
	numbers := make([]int, 0)
	fmt.Print("Provide sequence of numbers: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, v := range strings.Split(strings.ReplaceAll(input,"\r\n",""), " ") {
		value, err := strconv.Atoi(v)
		if err == nil {
			numbers = append(numbers, value)
		}
	}
	var wg sync.WaitGroup
	wg.Add(4)
	size := len(numbers)
	go bubbleSort(numbers[0:size/4], &wg)
	go bubbleSort(numbers[size/4:2*size/4], &wg)
	go bubbleSort(numbers[2*size/4:3*size/4], &wg)
	go bubbleSort(numbers[3*size/4:size], &wg)
	wg.Wait()
	fmt.Println(merge(
		merge(
		numbers[0:size/4],
		numbers[size/4:2*size/4]),
		merge(
			numbers[2*size/4:3*size/4],
			numbers[3*size/4:size])))
}
