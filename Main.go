// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {

	slice()

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

func slice() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	fmt.Println("----")
	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")
	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")
	// Double value range, ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("----")
	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}
