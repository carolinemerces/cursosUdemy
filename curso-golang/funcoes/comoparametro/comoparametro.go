package main

import "fmt"

//funcao comum nomeada
func multiplicacao(a, b int) int {
	return a*b
}

//funcao nomeada que recebe outra funcao nomada como parâmetro (que possui 2 parâmetros e um retorno) + 2 parâmetros e um retorno
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2) //a "funcao" está recebendo o parâmetro da funcao exec, ambos possuem o mesmo tipo de parâmetro (int)
}

func main(){
	resultado := exec(multiplicacao, 3, 4) //a variável resultado recebe como atribuicão a funcao exec, que recebe como parâmetro a funcao multiplicacao + os valores 3 e 4, conforme a funcao solicita valores de parâmetros
	fmt.Println(resultado) //será printada a variável resultado, que contém as diretrizes de todas as funcões relacionadas anteriormente -> a multiplicacão de 3 com 4 = 12
}