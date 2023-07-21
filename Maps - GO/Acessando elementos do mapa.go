package main

import "fmt"

func main() {

	fmt.Println("\n\t\tAcessando elementos do mapa:\n")

	/*SINTAXE:

	valor = mapaNome[chave]*/

	var a = make(map[string]string)
	a["Marca"] = "Ford"
	a["Modelo"] = "MUstang"
	a["Ano"] = "1964"

	fmt.Printf(a["Marca"])
	fmt.Println()
}
