package main

import "fmt"

//Peça ao usuário para digitar uma string e um caractere e retorne o número de
//ocorrências desse caractere na string.

func main() {
	var frase string
	var letra string
	total := 0

	fmt.Print("Qual a frase? ")
	fmt.Scan(&frase)

	fmt.Print("Qual a letra? ")
	fmt.Scan(&letra)

	for i := 0; i < len(frase); i++ {
		if string(frase[i]) == letra {
			total++
		}

	}
	fmt.Println("a letra apareceu", total, "vezes")

}
