package main

// Defer will execute all code e after fisnish will execute deferStack

import "fmt"

func main() {

	onedefer()
	stackdefer()

}

func onedefer() {

	defer fmt.Printf("Hello defer\n")

	fmt.Printf("Hello without defer\n")

}

func stackdefer() {

	for i := 0; i <= 10; i++ {

		defer fmt.Println(i)

	}

	fmt.Printf("Inicitialize\n")

}

func manydefers() {

	defer fmt.Println("Fisrt defer")

	defer fmt.Println("Second defer")

	defer fmt.Println("Third defer")

	for i := 0; i <= 10; i++ {

		defer onedefer()

	}

	fmt.Println("Initialize")

}
