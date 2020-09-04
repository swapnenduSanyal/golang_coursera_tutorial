package main

import "fmt"
import "strconv"
import "sort"

func main(){
	sli := make([]int,3,3);
	sli = sli[0:0];
	for c := "x"; c!="X";{
		fmt.Scan(&c);
		v,e := strconv.Atoi(c);
		if e==nil{
			sli = append(sli,v);
			sort.Ints(sli);
			fmt.Println(sli);
		}
	}
}

