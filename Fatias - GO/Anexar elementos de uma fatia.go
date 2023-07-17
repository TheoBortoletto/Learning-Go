package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tAnexar elementos de uma fatia: \n")

	/* SINTAXE:
	fatiaNome = append(fatiaNome, elemento1, elemento2, ...)*/

	myslice1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("Comprimento = %d\n", len(myslice1))
	fmt.Printf("Capacidade = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)

	fmt.Println()

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("Comprimento = %d\n", len(myslice1))
	fmt.Printf("Capacidade = %d\n", cap(myslice1))

}
