package main

import "fmt"

func main() {

	fmt.Println("\n\t\tIterando sobre mapas: \n")

	/*Este exemplo mostra como iterar sobre os elementos em um mapa.
	Observe a ordem dos elementos na saída.*/

	a := map[string]int{"Um": 1, "Dois": 2, "Três": 3, "Quatro": 4}

	for i, j := range a {
		fmt.Printf("%v : %v,", i, j)
	}
	fmt.Println()
}
