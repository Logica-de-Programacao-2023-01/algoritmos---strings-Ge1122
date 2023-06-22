package main

import "fmt"

//Peça ao usuário para digitar uma string e retorne o número de caracteres nessa string.

func main() {
	var x string
	fmt.Print("Escreva uma palavra: ")
	fmt.Scan(&x)
	s := x
	fmt.Println(len(s))
}
