package main

import "fmt"

func main() {

	fmt.Println("\n\t\tCriando um mapa vazio:\n")

	/*Existem duas maneiras de criar um mapa vazio. Uma é
	usando a função make() e a outra é usando a seguinte sintaxe:

	var a map[TipoChave]TipoValor

	OBS: A função make() é o caminho certo para
	criar um mapa vazio. Se você criar um mapa
	vazio de uma maneira diferente e escrever
	nele, isso causará um pânico no tempo de
	execução.
	*/

	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil) // -> EM GO, 'nil' é um valor nulo para ponteiros, maps, slices etc...
	fmt.Println(b == nil)

}
