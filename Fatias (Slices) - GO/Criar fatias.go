package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tCriar uma fatia a partir de uma matriz: \n")

	/*
		SINTAXE:

		var mtx = [comprimento]tipo de dado{valores} -> MATRIZ
		myslice := mtx[início:final] -> FATIA feita a partir da matriz acima
	*/

	mtx := [6]int{10, 11, 12, 13, 14, 15}
	myslice := mtx[2:4] // Pegará os elementos da posição 2 até o 4, sem contar com a posição 4.

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("Comprimento = %d\n", len(myslice))
	fmt.Printf("Capacidade = %d\n", cap(myslice))

	/*No exemplo acima mysliceé uma fatia com comprimento 2. É feito de arr1qual é um array com comprimento 6.

	A fatia começa no segundo elemento da matriz que tem valor 12. A fatia pode crescer até o final da matriz.
	Isso significa que a capacidade da fatia é 4.

	Se mysliceiniciado do elemento 0, a capacidade da fatia seria 6.*/

}
