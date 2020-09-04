package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		name    string
		address string
	}
	m := make(map[string]string)
	//Set key/value pairs using typical name[key] = val syntax.
	fmt.Print("Enter Name : ")
	var name string
	fmt.Scanln(&name)
	m["name"] = name
	fmt.Print("Enter Address : ")
	var address string
	fmt.Scanln(&address)
	m["address"] = address

	per, err := json.Marshal(m)
	if err == nil {
		fmt.Println("Final JSON Byte Array Output :", per)
		fmt.Println("Final JSON Output :", string(per))
	}
}
