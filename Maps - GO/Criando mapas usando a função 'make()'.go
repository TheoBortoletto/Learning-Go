package main

import "fmt"

func main() {

	fmt.Println("\n\t\tCriando mapas usando a função 'make()': \n")

	/*SINTAXE:

	var a = make(map[TipoChave]TipoValor)
	b := make(map[TipoChave]TipoValor)
	*/

	var a = make(map[string]string) // -> MAPA VAZIO
	a["Marca"] = "Ford"
	a["Modelo"] = "Mustang"
	a["Ano"] = "1964" // -> MAPA NÃO ESTÁ MAIS VAZIO

	b := make(map[string]int)
	b["Prova1"] = 4
	b["Prova2"] = 8
	b["Prova3"] = 7
	b["Prova4"] = 2

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("a\t%v\n", b)
}
