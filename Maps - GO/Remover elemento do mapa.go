package main

import "fmt"

func main() {

	fmt.Println("\n\t\tRemover elementos do mapa: \n")

	/*A remoção de elementos é feita usando a função 'delete()'.

	SINTAXE:

	delete(mapaNome, chave)*/

	var a = make(map[string]string)
	a["Marca"] = "Ford"
	a["Modelo"] = "Mustang"
	a["Ano"] = "1964"

	fmt.Println("Mapa com todos os elementos: ", a, "\n")

	delete(a, "Ano") // DELETANDO O ELEMENTO 'ANO' DO MAPA 'a'

	fmt.Println("Mapa com o elemento 'Ano' deletado: ", a)

}
