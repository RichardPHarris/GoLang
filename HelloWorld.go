// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {

	fizzBuzz()

}

func Func2_1() {
	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, y)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("Result: %v, type of %T\n", mean, mean)
}

func Func2_2If() {

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

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more then half of b")
	}

}

func Func2_2Switch() {

	x := 2

	switch x {

	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Many")
	}

	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}

}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func fizzBuzz() {

	for i := 0.0; i < 21; i++ {

		if i/3 == float64(int64(i/3)) {
			fmt.Println("Fizz")
		} else if i/5 == float64(int64(i/5)) {
			fmt.Println("Buzz")
		} else if i/5 == float64(int64(i/5)) && i/3 == float64(int64(i/3)) {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
