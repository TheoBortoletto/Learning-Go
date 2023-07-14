package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tVerbos de Formatação para 'Printf()': \n\n\tVerbos de Formatação Geral:\n")

	/*Os seguintes verbos podem ser usados ​​com todos os tipos de dados:

	Verb			Description
	%v	->	Imprime o valor no formato padrão
	%#v	->	Imprime o valor na sintaxe 'Go'
	%T	->	Imprime o tipo do valor (int, string...)
	%%	->	Imprime o sinal %*/

	var i = 20.7
	var txt = "Olááá a todos!!!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n\n", txt)

	fmt.Println("\tVerbos de formatção inteira: \n")

	/*Os seguintes verbos podem ser usados ​​com o tipo de dados inteiro:

	Verb: 	Description:
	%b	 ->	 Base 2
	%d	 ->	 Base 10
	%+d	 ->	 Base 10 e sempre mostra o sinal
	%O	 ->	 Base 8, com o começo com '0o'
	%x	 ->	 Base 16, minúsculo
	%X	 ->	 Base 16, maiúsculo
	%#x	 ->	 Base 16, começando com '0x'
	%4d	 ->	 Pad with spaces (width 4, right justified)
	%-4d ->	 Pad with spaces (width 4, left justified)
	%04d ->	 Pad with zeroes (width 4)*/

	var num = 15

	fmt.Printf("%b\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%+d\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%O\n", num)
	fmt.Printf("%x\n", num)
	fmt.Printf("%X\n", num)
	fmt.Printf("%4d\n", num)
	fmt.Printf("%-4d\n", num)
	fmt.Printf("%04d\n", num)

}
