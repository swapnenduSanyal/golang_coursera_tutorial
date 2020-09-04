package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {

	type Name struct {
		fname string
		lname string
	}
	names := make([]Name,0)
	var filename string
	fmt.Scan(&filename)
	data, _ := ioutil.ReadFile(filename)
	split := strings.Split(string(data),"\n")
	for _,v := range split{
		split2:= strings.Split(strings.Trim(v," ")," ")
		if len(split2) == 2{
			names = append(names,Name{split2[0],split2[1]})
		}
	}
	for _,v := range names{
		fmt.Println(v)
	}
}
