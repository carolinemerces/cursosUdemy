package main

import "fmt"

//funcao recursividade mais simples

func fatorial(n uint) uint { //utilizando uint não se aceita valores inteiros negativos, que deixam nossa aplicacao mais complexa
	switch { //com valor verdadeiro
	case n == 0: //o fatorial de 1 e 0, representam o resultado 1
		return 1 
	default: //nos demais casos, será feito o fatorial do número n, ou seja, vai fazer uma chamada recursiva chamando a funcao fatorial com n-1
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5) //criar uma variável resultado, que recebe a funcão fatorial com o valor 5
	fmt.Println(resultado)
	}

	//neste caso não permite aplicar um valor negativo