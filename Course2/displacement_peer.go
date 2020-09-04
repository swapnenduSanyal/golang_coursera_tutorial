package main

import "fmt"

func GenDisplaceFn(acc, vel, dis float64) func(float64) float64{
	fu := func(time float64) float64{
		return (acc*(time*time)/2) + (vel*time) + dis
	}
	return fu
}

func main() {
	var (
		acc float64
		vel float64
		dis float64
		time float64
		cal func(float64) float64
	)
	fmt.Printf("Enter a value for acceleration: ")
	_,_ = fmt.Scan(&acc)
	fmt.Printf("Enter a value for Initial velocity: ")
	_,_= fmt.Scan(&vel)
	fmt.Printf("Enter a value for Initial Displacement: ")
	_,_= fmt.Scan(&dis)
	fmt.Printf("Enter a value for Time: ")
	_,_= fmt.Scan(&time)
	cal = GenDisplaceFn(acc, vel, dis)
	fmt.Println(cal(time))
}