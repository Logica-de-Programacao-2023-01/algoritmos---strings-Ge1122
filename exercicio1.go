package main

import "fmt"

func main() {
	var x, y string
	fmt.Print("Escreva uma palavra: ")
	fmt.Scan(&x)
	fmt.Print("Escreva outra frase: ")
	fmt.Scan(&y)

	s1 := x
	s2 := y
	s3 := s1 + ", " + s2
	fmt.Println(s3)

}
