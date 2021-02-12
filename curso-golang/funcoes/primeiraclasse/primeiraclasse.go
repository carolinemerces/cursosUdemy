package main

import "fmt"

//funcao anônima - a partir da variável soma, que recebe essa funcao sem nome, é possível executar a funcao
var soma = func (a, b int) int {
	return a+b
}

func main(){
	fmt.Println(soma(2, 3))

//dentro de uma funcão é possível ter outra funcão ou passar como parâmetro de uma funcão, outra funcão
	sub := func (a, b int) int {
		return a-b
	}
	fmt.Println(sub(5, 3))
}
