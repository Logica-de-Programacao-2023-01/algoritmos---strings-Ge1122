package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y, z string
	fmt.Println("Digite uma string: ")
	fmt.Scan(&x)

	fmt.Println("Digite qual letra sera subtituida: ")
	fmt.Scan(&y)

	fmt.Println("Digite a letra substituta: ")
	fmt.Scan(&z)

	resultado := strings.ReplaceAll(x, y, z)

	fmt.Println("A string modificada é: %s\n", resultado)
}
