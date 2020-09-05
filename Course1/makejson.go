package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var name, address string
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Enter address")
	fmt.Scan(&address)
	data, _ := json.Marshal(map[string]string{
		"name":name,
		"address":address,
	})
	fmt.Println(data)
	fmt.Println(string(data))
}
