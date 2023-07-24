package main

import "fmt"

func main() {

	fmt.Println("\n\t\tMapas são referências: \n")

	/*Os mapas são referências a tabelas hash.

	Se duas variáveis ​​de mapa se referem à mesma tabela de hash,
	alterar o conteúdo de uma variável afeta o conteúdo da outra.*/

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}
