package main

import (
	"fmt"

	"joaoc.com/learning/guessgaming"
)

func main() {

	greetins := fmt.Sprintf("Hello, %s", "Joao")

	fmt.Println(greetins)

	guessgaming.GameStart(100)

}
