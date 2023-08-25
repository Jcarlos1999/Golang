package main

import "fmt"

var x, y int = 23, 33

func main() {

	p := &x
	fmt.Printf("Memory allocate of x: %d\n", p)
	fmt.Printf("Value of x using pointer: %d\n\n", *p)

	*p = 999

	fmt.Printf("Memory allocate of x: %d\n", p)
	fmt.Printf("Change x value to %d by pointer\n\n", x)

	p = &y
	fmt.Printf("Memory allocate of y: %d\n", p)
	fmt.Printf("Value of y using pointer: %d\n\n", *p)

	*p = 666
	fmt.Printf("Memory allocate of y: %d\n", p)
	fmt.Printf("Change y value to %d by pointer\n", y)

}
