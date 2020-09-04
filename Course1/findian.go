package main

import "fmt"
import "strings"

func main(){
	var x string;
	fmt.Scan(&x);
	x = strings.ToLower(x);
	if x[0] == 'i' && len(x)>0 && x[len(x)-1] == 'n' && strings.ContainsRune(x,'a'){
		fmt.Println("Found!");
	} else{
		fmt.Println("Not Found!");
	}
}

