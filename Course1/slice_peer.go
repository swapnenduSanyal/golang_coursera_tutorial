package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input:=make([]int,0)
	var lineInput string
	for i:=0;;i++ {
		fmt.Println("Enter Numbers and if you want to exit, Enter 'X'")
		fmt.Scanf("%s\n",&lineInput)
		if lineInput[0] =='X'{
			break
		}else {
			temp, _ :=strconv.Atoi(lineInput)
			input=append(input, temp)
		}
		sort.Ints(input)
		fmt.Println(input)
	}
	sort.Ints(input)
	fmt.Println(input)

}