package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tIntrodução a fatias: \n")

	/* Fatias são semelhantes a matrizes, mas são mais poderosas e flexíveis.

	Assim como os arrays, os slices também são usados ​​para armazenar vários valores do mesmo tipo em uma única variável.

	No entanto, ao contrário das matrizes, o comprimento de uma fatia pode aumentar e diminuir conforme você desejar.

	Em Go, existem várias maneiras de criar uma fatia:

	Usando o formato [] tipo de dados { valores }
	Criar uma fatia de uma matriz
	Usando a função make() */

	fmt.Println("\tSintaxe:")

	// myslice := []int{}
	// O código acima declara uma fatia vazia de comprimento 0 e capacidade 0.

	// Para inicializar a fatia durante a declaração, use isto:

	// myslice := []int{1, 2, 3}
	// O código acima declara uma fatia de inteiros de comprimento 3 e também a capacidade de 3.

	/* Em Go, existem duas funções que podem ser usadas para retornar o comprimento e a capacidade de uma fatia:

	len()função - retorna o comprimento da fatia (o número de elementos na fatia)
	cap()função - retorna a capacidade da fatia (o número de elementos que a fatia pode aumentar ou diminuir) */

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1, "\n")

	myslice2 := []string{"Python", "é", "melhor", "que", "Go"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2, "\n")

	/* No exemplo acima, vemos que na primeira fatia (myslice1),
	os elementos reais não são especificados, portanto, tanto o
	comprimento quanto a capacidade da fatia serão zero. Na segunda fatia
	(myslice2), os elementos são especificados e tanto o comprimento quanto
	a capacidade são iguais ao número de elementos reais especificados. */

}
