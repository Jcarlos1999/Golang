package guessgaming

import (
	"fmt"
	"math/rand"
)

var userInput int
var secretNumber int

func GameStart(maxRange int) {

	if maxRange == 0 {
		secretNumber = 100
	} else {
		secretNumber = rand.Intn(maxRange)

	}
	fmt.Println("Escolha um numero")
	fmt.Scanln(&userInput)
	for {
		if userInput == secretNumber {
			fmt.Println("Parabens voce acertou o numero secreto")
			break
		} else if userInput > secretNumber {
			fmt.Println("Seu numero é maior que o numero secreto! Escolha outro numero")
			fmt.Scanln(&userInput)
		} else {
			fmt.Println("Seu numero é menor que o numero secreto! Escolha outro numero")
			fmt.Scanln(&userInput)
		}

	}

}
