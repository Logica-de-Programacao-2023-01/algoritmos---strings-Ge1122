package main

import (
	"fmt"
	"strings"
)

//Peça ao usuário para digitar uma string e um número n e retorne a mesma string com as
//n primeiras letras em maiúsculo.

func main() {
	var frase string
	var x int
	var total string

	fmt.Print("Qual é a frase? ")
	fmt.Scan(&frase)

	fmt.Print("quantas letras você quer maiuscula? ")
	fmt.Scan(&x)

	for i := 0; i < x; i++ {
		fmt.Print(strings.ToUpper(string(frase[i])))
	}
	for a := 0; a > x; a++ {
		fmt.Print(strings.ToLower(string(frase[a])))
	}
	total = (frase + frase)
	fmt.Print(total)

}
