package main

import "fmt"

func Func2_1() {
	var x int
	var y int

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, y)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean int
	mean = (x + y) / 2
	fmt.Printf("Result: %v, type of %T\n", mean, mean)

}
