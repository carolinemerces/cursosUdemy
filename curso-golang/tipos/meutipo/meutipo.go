package main

import "fmt"

//tipos personalizados a partir de um tipo base, utilizando como receiver, com um apelido e  associar novos métodos seus
//exemplo nota para conceito
//uso do if else ou switch

type nota float64 //tipo nosso definido/personalizado

//funcao criada a partir do tipo float64 e identificador nota
//descobrir se a nota está dentro de um conjunto de notas (está entre inicio e fim?)
func (n nota) entre(inicio, fim float64) bool {
	//forma encapsulada dessa funcao, utilizando && na funcao, ao invés de utilizar abaixo no if
	return float64(n) >= inicio && float64(n) <= fim //se estiver dentro dessa condicao retornará verdadeiro
}

//essa funcao recebe como parâmetro o tipo float64 e identificador nota, retornando uma string
//desafio refatorar utilizando switch case default
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0){
		return "A"
	} else if n.entre(7.0, 8.99){
		return "B"
	} else if n.entre(5.0, 7.99){
		return "C"
	} else if n.entre(3.0, 4.99){
		return "D"
	} else {
		return "E"
	}
}

func main(){
	fmt.Println(notaParaConceito(9.6))
	fmt.Println(notaParaConceito(6.6))
	fmt.Printf("Sua nota é %2.f e você ficou com %s", 2.6, notaParaConceito(2.6))

}