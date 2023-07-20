package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\t\riar uma fatia com a função 'make()': \n")

	/* A funçãa 'make()' também pode ser usada para criar uma fatia

	SINTAXE:

	nomeDaFatia := make([]tipo, comprimento, capacidade)

	OBS.: Caso o parâmetro capacidade não seja definido, será igual ao comprimento .*/

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("Comprimento = %d\n", len(myslice1))
	fmt.Printf("Capacidade = %d\n", cap(myslice1))

	/* Agora uma fatia sem a capacidade: */

	myslice2 := make([]int, 5)
	fmt.Printf("myslice = %v\n", myslice2)
	fmt.Printf("Comprimento = %d\n", len(myslice2))
	fmt.Printf("Capacidade = %d\n", cap(myslice2))

}
