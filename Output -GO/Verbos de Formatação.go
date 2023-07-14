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
	var txt_ = "Olááá a todos!!!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt_)
	fmt.Printf("%#v\n", txt_)
	fmt.Printf("%T\n\n", txt_)

	fmt.Println("\tVerbos de formatção inteira (int): \n")

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
	fmt.Printf("%04d\n\n", num)

	fmt.Println("\tVerbos de formatação de string (string): \n")

	/*Os seguintes verbos podem ser usados ​​com o tipo de dados string:

	Verb:	Description:
	%s	 ->	 Imprime o valor como string simples
	%q	 ->	 Prints the value as a double-quoted string
	%8s	 ->	 Imprime o valor como string simples (largura 8, direita justified)
	%-8s ->	 Imprime o valor como string simples (largura 8, esquerda justified)
	%x	 ->	 Imprime o valor hexa decimal sem espaço
	% x	 ->	 Imprime o valor hexa decimal com espaços*/

	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n\n", txt)

	fmt.Println("\tVerbos de formatação booleana (bool): \n")

	/*O verbo a seguir pode ser usado com o tipo de dados booleano:

	Verb:	Description:
	%t	->	IMprime o valor se é veerdadeiro ou falso (o mesmo se usar %v)*/

	var a = true
	var b = false

	fmt.Printf("%t\n", a)
	fmt.Printf("%t\n\n", b)

	fmt.Println("\tVerbos de formatação flutuante (float): \n")

}
