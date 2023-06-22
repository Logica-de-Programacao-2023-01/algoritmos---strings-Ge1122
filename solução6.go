package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	palavras := strings.Fields(x)
	qntpalavras := len(palavras)

	fmt.Printf("A string contém %d palavra(s).\n", qntpalavras)
}
