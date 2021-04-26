// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!!")
	Func2_1()

}

func Func2_1() {
	var x float64
	var y float64

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, y)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2.0
	fmt.Printf("Result: %v, type of %T\n", mean, mean)
}
