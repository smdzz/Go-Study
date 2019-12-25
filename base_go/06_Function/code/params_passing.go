package main

import (
	"fmt"
)

func main() {
	x := 3
	fmt.Println("x = ", x, "&x = ", &x)

	y := add(x)
	fmt.Println("x = ", x, "y = ", y)

	z := addP(&x)
	fmt.Println("x = ", x, "z = ", z)
	fmt.Println("&x = ", &x)

}

func add(x int) int {
	x++
	return x
}

func addP(x *int) int {
	*x++
	return *x
}