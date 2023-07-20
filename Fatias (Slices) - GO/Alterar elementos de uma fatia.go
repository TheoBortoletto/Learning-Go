package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tAlterar elementos de uma fatia: \n")

	valores := []int{10, 20, 30}
	valores[2] = 50 // Alterou o terceiro elemento da faita

	fmt.Println(valores[0])
	fmt.Println(valores[2])
}
