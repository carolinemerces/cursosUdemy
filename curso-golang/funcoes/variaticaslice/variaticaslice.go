package main

import "fmt"

//recebe um conjunto variaveis de parametros do tipo string e a funcao irá imprimir essa lista de string
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados { //uso do fro range para varrer os parâmetros e adicioná-los na variável aprovado, dessa vez também utilizando o índice para numerar essa lista
		fmt.Printf("%d) %s\n", i+1, aprovado) //printará na tela: 1) fulano; 2) funalo...
	}
}

func main(){
	aprovados := []string {"Maria", "João", "André"} //criacão de um slice de aprovados, slice pois não há definicão do tamanho dele
	imprimirAprovados(aprovados...)//para passar um slice para uma funcao variatica é só colocar a variável definida com os ... / e assim os elementos da variável do slice serão espalhados como se fossem parâmetros para essa execucão

	//importante: neste contexto de funcão variática não é possível utilizar um array ou os ... dentro do slice, pois não funciona
}