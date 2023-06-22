
package main


import (
	"fmt"
	"strconv"
)
func main() {
	var x string
	fmt.Print("Digite uma sequência numérica: ")
	fmt.Scan(&x)
  
	previonum := -1
	isIncreasing := true
  
	for _, char := range x {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Entrada inválida.")
			return
		}
		if num <= previonum {
			isIncreasing = false
			break
		}
		previonum = num
	}

	if isIncreasing {
		fmt.Println("A sequência é numérica crescente.")
	} else {
		fmt.Println("A sequência não é numérica crescente.")
	}
}
