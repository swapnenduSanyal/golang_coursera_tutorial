package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	maxLength = 20
)

func main() {

	type Student struct {
		fname string
		lname string
	}
	var a Student
	sli := make([]Student, 0)

	fmt.Print("Enter The File Name: ")
	var fileName string
	fmt.Scanln(&fileName)

	fileName = strings.Trim(fileName, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), " ")

			a.fname = s[0]
			a.lname = s[1]
			sli = append(sli, a)
		}
	}
	file.Close()

	for _, v := range sli {
		fmt.Println(v.fname, v.lname)
	}

}
