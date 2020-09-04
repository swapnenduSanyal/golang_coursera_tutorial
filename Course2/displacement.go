package main

import "fmt"

func GenDisplaceF(s0,u,a float64) func(float64) float64 {
	return func(t float64) float64{
		return s0+u*t+0.5*a*t*t
	}
}

func main(){
	var s0,u,a,t float64
	fmt.Print("Provide acceleration, velocity and displacement separated by spaces :");
	fmt.Scanf("%f %f %f",&a,&u,&s0)
	f := GenDisplaceF(s0,u,a)
	fmt.Print("Enter time :");
	fmt.Scan(&t)
	fmt.Println("Displacement :",f(t))
}

