package main

import "fmt"

func main() {

	fmt.Println("\n\t\tAtualizando e adicionando elementos no mapa: \n")

	/*Atualizar ou adicionar um elemento é feito por:

	SINTAXE:

	mapaNome[chave] = valor*/

	var a = make(map[string]string)
	a["Marca"] = "Ford"
	a["Modelo"] = "Mustang"
	a["Ano"] = "1964"

	fmt.Println("Mapa SEM ALTERAÇÃO: ", a)

	a["Ano"] = "1970"     // ATUALIZANDO UM ELEMENTO
	a["Cor"] = "Vermelho" // ADICIONANDO UM ELEMENTO

}
