package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tAnexar uma fatia a outra fatia: \n")

	/*
		Para anexar todos os elementos de uma fatia a outra fatia, use a função 'append()'':

		fatia1 = append(fatia1, fati2a)

		Observação: O '...' após fatia2 é necessário ao anexar os elementos de uma fatia a outra.
	*/

	fatia1 := []int{1, 2, 3}
	fatia2 := []int{4, 5, 6}
	fatia3 := append(fatia1, fatia2...) // Anexando duas fatias em uma só

	fmt.Println("fatia1: ", fatia1)
	fmt.Println("fatia2: ", fatia2, "\n")

	fmt.Printf("fatia3 = %v\n", fatia3)
	fmt.Printf("Comprimento = %d\n", len(fatia3))
	fmt.Printf("Capacidade = %d\n", cap(fatia3))
}
