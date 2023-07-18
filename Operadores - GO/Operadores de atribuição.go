package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n\t\tOperadores de atribuição: \n")

	/*Operadores de atribuição são usados ​​para atribuir valores a variáveis.

	No exemplo abaixo, utilizamos o operador de atribuição= ( ) para atribuir o
	valor 10 a uma variável chamada x :*/

	var x = 10

	fmt.Println(x, "\n")

	fmt.Println("Operador de atribuição de adição: \n")

	/*O operador de atribuição de adição ( +=) adiciona um valor a uma variável:*/

	var num = 10
	num += 5 // -> 10 + 5 = 15
	fmt.Println(num, "\n")

}

/*

Uma lista de todos os operadores de atribuição:

Operador:	Exemplo:	  Mesmo que:

=	    |   x = 5	|	x = 5
+=	    |   x += 3  |	x = x + 3
-=	    |   x -= 3  |	x = x - 3
*=	    |   x *= 3  |	x = x * 3
/=	    |   x /= 3  |	x = x / 3
%=	    |   x %= 3  |	x = x % 3
&=	    |   x &= 3  |	x = x & 3
|=	    |   x |= 3  |	x = x | 3
^=	    |   x ^= 3  |	x = x ^ 3
>>=	    |   x >>= 3 |	x = x >> 3
<<=	    |   x <<= 3 |	x = x << 3

Veja os exemplos em: https://www.w3schools.com/go/go_assignment_operators.php

*/
