package main

import "fmt"

func swap(numbers []int, i int) {
	if numbers[i]>numbers[i+1]{
		temp := numbers[i]
		numbers[i] = numbers[i+1]
		numbers[i+1] = temp
	}
}

func BubbleSort(sli []int){
	for i:= 0; i<len(sli)-1; i++{
		swap(sli,i)
	}
	if len(sli)>1{
		BubbleSort(sli[0:len(sli)-1])
	}
}

func main(){
	numbers := make([]int,10)
	fmt.Print("Provide a sequence of upto 10 numbers: ");
	size,_ := fmt.Scanln(&numbers[0],
		&numbers[1],
		&numbers[2],
		&numbers[3],
		&numbers[4],
		&numbers[5],
		&numbers[6],
		&numbers[7],
		&numbers[8],
		&numbers[9])
	BubbleSort(numbers[0:size])
	fmt.Println(numbers[0:size])
}

