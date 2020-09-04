package main

import (
	"fmt"
	"strings"
)

func main() {
	var i string
	fmt.Scan(&i)
	// fmt.Print(strings.ToLower(i))
	// fmt.Print(string(i[len(i)-1]))
	i = strings.ToLower(i)
	if strings.HasPrefix(i, "i") && strings.HasSuffix(i, "n") && strings.Contains(i, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
