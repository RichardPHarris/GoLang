// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	Func2_2()

}

func Func2_1() {
	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, y)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("Result: %v, type of %T\n", mean, mean)
}

func Func2_2() {

	x := 10

	if x > 5 {
		fmt.Println("x is big")

	}
	if x > 100 {
		fmt.Println("x is very big")

	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}
}
