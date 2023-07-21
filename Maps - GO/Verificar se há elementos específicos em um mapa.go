package main

import "fmt"

func main() {

	fmt.Println("\n\t\tVerificar se há elementos específicos em um mapa: \n")

	/*Você pode verificar se uma determinada chave existe em um mapa usando:

	SINTAXE:

	val, ok := nomeMapa[chave]

	Se você deseja apenas verificar a existência de uma
	determinada chave, pode usar o identificador em branco ( _ ) no lugar de 'val'.*/

	var a = map[string]string{"Marca": "Ford", "Modelo:": "Mustang", "Ano": "1964", "Dia": ""}

	val1, ok1 := a["Marca"] // CHECANDO A CHAVE EXISTENTE E O SEU VALOR
	val2, ok2 := a["Cor"]   // CHECANDO A CHAVE INEXISTENTE E O SEU VALOR
	val3, ok3 := a["Dia"]   // CHECANDO A CHAVE EXISTENTE E O SEU VALOR
	_, ok4 := a["Modelo"]   // VERIFICANDO APENAS A CHAVE EXISTENTE E NÃO O SEU VALOR

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

}
