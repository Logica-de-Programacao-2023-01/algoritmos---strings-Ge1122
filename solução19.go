package main

import "fmt"



func main() {
	var x string
	fmt.Println("Digite uma string:")
	fmt.Scan(&x)

	inverter := ""
	for i := len(x) - 1; i >= 0; i-- {
		inverter += string(x[i])
	}
	fmt.Println("String invertida:", inverter)
}
