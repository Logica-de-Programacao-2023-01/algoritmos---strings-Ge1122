package main

//Peça ao usuário para digitar uma string e retorne a mesma string com as letras em
//maiúsculo.
import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	fmt.Println("Qual a frase? ")
	fmt.Scan(&frase)
	fmt.Println(strings.ToUpper(frase))

}
