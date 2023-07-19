package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tLoops aninhados: \n")

	/*É possível colocar um loop dentro de outro loop.
	Aqui, o "loop interno" será executado uma vez para cada iteração do "loop externo":*/

	adjetivos := [2]string{"Grande", "Bom"}
	frutas := [3]string{"Maçã", "Banana", "Laranja"}

	for i := 0; i < len(adjetivos); i++ {
		for j := 0; j < len(frutas); j++ {
			fmt.Println(adjetivos[i], frutas[j])
		}
	}
}
