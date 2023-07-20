package main

import "fmt"

func main() {

	fmt.Println("\n\t\tEstruturas (Structs):\n")

	/*Um struct (abreviação de structure) é usado para criar uma
	coleção de membros de diferentes tipos de dados em uma única variável.

	Enquanto arrays são usados ​​para armazenar vários valores
	do mesmo tipo de dados em uma única variável, structs são
	usados ​​para armazenar vários valores de diferentes tipos
	de dados em uma única variável.

	Uma struct pode ser útil para agrupar dados para criar registros.*/

	/*SINTAXE:

		type struct_name struct {
	  		member1 datatype;
	 	 	member2 datatype;
	  		member3 datatype;
	  ...
	}
	*/

	type Pessoa struct {
		nome    string
		idade   int
		job     string
		salario int
	}

}
