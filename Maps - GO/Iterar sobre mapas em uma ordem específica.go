package main

import "fmt"

func main() {

	fmt.Println("\n\t\tIterar sobre mapas em uma ordem específica: \n")

	/*Os mapas são estruturas de dados não ordenadas.
	Se você precisar iterar em um mapa em uma ordem específica,
	deverá ter uma estrutura de dados separada que especifique essa ordem.*/

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

}
