package main

import (
	"fmt"
	"unicode"
)


func main() {
	var x string
	fmt.Println("Digite uma string:")
	fmt.Scan(&x)

	numero := false
	for _, char := range x  {
		if unicode.IsNumber(char) {
			numero = true
			break
		}
	}
	if numero {
		fmt.Println("A string contém pelo menos um número")
	} else {
		fmt.Println("A string não contém nenhum número")
	}
}
