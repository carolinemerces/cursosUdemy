package main

import "fmt"

func main () {
	numeros := [...] int {1, 2, 3, 4, 5} //com os ... é um array, onde o compilador irá contar / sem ... é um slice

	for i, numero := range numeros { //a partir do range associado com o array, me retorna o índice e o elemento do array que está sendo percorrido dentro do array

		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros { //para ignorar o índice e pegar o que está dentro do array, utiliza-se _
		fmt.Println(num)
	}
}