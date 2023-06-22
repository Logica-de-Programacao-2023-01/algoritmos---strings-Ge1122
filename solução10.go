package main

import (
	"fmt"
	"sort"
	"strings"
)

//Escreva um programa que receba duas strings
//e verifique se elas são anagramas.
//Informe ao usuário se são ou não.

func main() {
	var x, y string
	fmt.Println("Digite a primeira string:")
	fmt.Scan(&x)
	fmt.Println("Digite a segunda string:")
	fmt.Scan(&y)
 
  y = strings.ToLower(y)
	x = strings.ToLower(x)
	

	if len(y) != len(x) {
		fmt.Println("As strings não são anagramas")
		return
	}
	runes1 := []rune(x)
	runes2 := []rune(y)
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	if string(runes1) == string(runes2) {
		fmt.Println("As strings são anagramas")
	} else {
		fmt.Println("As strings não são anagramas")
	}
}
