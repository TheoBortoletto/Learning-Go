package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tAlterar o comprimento de uma fatia\n")

	// Ao contrário dos arrays (matrizes), é possível alterar o comprimento de uma fatia.

	mtx1 := [6]int{9, 10, 11, 12, 13, 14} // MATRIZ
	fatia1 := mtx1[1:5]                   // FATIA DA MATRIZ

	fmt.Printf("fatia1 = %v\n", fatia1)
	fmt.Printf("Comprimento = %d\n", len(fatia1))
	fmt.Printf("Capacidade = %d\n", cap(fatia1))
	fmt.Println()

	fatia1 = mtx1[1:3] // Alterando a matriz cortando novamente a matriz

	fmt.Printf("fatia1 = %v\n", fatia1)
	fmt.Printf("Comprimento = %d\n", len(fatia1))
	fmt.Printf("Capacidade = %d\n", cap(fatia1))
	fmt.Println()

	fatia1 = append(fatia1, 20, 21, 22, 23) // Mudando o comprimento anexando mais itens

	fmt.Printf("fatia1 = %v\n", fatia1)
	fmt.Printf("Comprimento = %d\n", len(fatia1))
	fmt.Printf("Capacidade = %d\n", cap(fatia1))s

}
