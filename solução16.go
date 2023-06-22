package main

import "fmt"



func main() {
	var x, y string
	fmt.Println("Digite a primeira string:")
	fmt.Scan(&x)

	fmt.Println("Digite a segunda string:")
	fmt.Scan(&y)

	if len(x) < len(y) {
		fmt.Println("Erro: a segunda string é maior que a primeira")
	} else {
		found := false
    
		for i := 0; i <= len(x)-len(y); i++ {
			if x[i:i+len(y)] == y {
				found = true
				break
			}
		}
    
		if found {
			fmt.Println("A segunda string é uma substring da primeira")
		} else {
			fmt.Println("A segunda string não é uma substring da primeira")
		}
	}
}
