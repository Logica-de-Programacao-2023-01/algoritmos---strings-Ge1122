package main

import "fmt"


func main() {
	var x string
	fmt.Println("Digite uma string:")
	fmt.Scan(&x)

	runes := []rune(x)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	output := string(runes)
	fmt.Printf("A string invertida é: %s\n", output)
}
