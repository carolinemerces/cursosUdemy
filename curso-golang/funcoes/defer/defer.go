package main

import "fmt"

//funcao defer -> funcao para atrasar a execucao da sentenca de código, até momento antes de retornar, do método sair chamar o return
//coloca "defer" antes de uma sentenca de código
//uso comum para fechar algum de tipo de recurso aberto

func obterValorAprovado(num int) int {
	//com defer segundo a ser executado
	//sem defer primeiro a ser executado 
	defer fmt.Println("Fim!") //essa sentenca será atrasada e executada depois do if e um pouco antes de terminar a funcao, antes do return
	if num > 5000 {
		//com defer primeiro a ser executado
		//sem defer segundo a ser executado
		fmt.Println("Valor alto...") 
	} else {
		fmt.Println("Valor baixo...")	
	}
	//com defer terceiro a ser executado
	//sem defer terceiro a ser executado
	return num 
}

func main(){
	fmt.Println(obterValorAprovado(6000))
}
