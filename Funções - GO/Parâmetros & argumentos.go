package main

import (
	"fmt"
)

func nomeFamilia(fname string) {
	fmt.Println("Olá", fname, "Bortoletto")
}

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {

	fmt.Println("\n\t\tParâmetros e argumentos:\n")

	/*O exemplo a seguir tem uma função com um parâmetro
	(fname) do tipo string. Quando a função familyName() é chamada,
	também passamos um nome (por exemplo, Liam),
	e o nome é usado dentro da função, que gera vários
	nomes diferentes, mas um sobrenome igual:*/

	nomeFamilia("Théo")
	nomeFamilia("Fred")

	fmt.Println("\nParâmetros múltiplos:\n")

	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

}
