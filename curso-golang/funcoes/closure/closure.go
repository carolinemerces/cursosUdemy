package main

import "fmt"

//Em Go, uma funcão também é um closure (fechamento / algo que envolve outra), quando há uma funcão interna em outra funcão / a funcao interna tem memória (tenha ciência) do local a qual foi declarada, ou seja, em que contexto léxico (local físico) ela foi declarada, funcionando de forma local.
//Nesse contexto a funcão externa funciona como um fechamaneto, agrupando o que está dentro, de forma que existam vários contextos diferentes declarados
//exemplo: uma variável declarada dentro de uma funcão não pode ser utilizada fora / então é possível declarar duas variáveis com mesmo nome, mas em contextos diferentes
func closure() func() { //funcao closure, sem parâmetros, que returna uma funcão anônima sem parâmetros 
	x := 10 //variável x no primeiro contexto
	var funcao = func() { //a variável funcao recebe a funcao anônima sem parâmetros func, que irá imprimir a variável x
		fmt.Println(x) 
	}
	return funcao //retorno da funcao (variável) criada como resposta da funcao maior/externa (closure)
}

func main() {
	x := 20 //variável x no segundo contexto
	fmt.Println(x)

	//criacão da variável imprimeX que recebe o resultado da chamada da funcão closure (=10). Consegue acessar devido a memória de onde e qual escopo foi definida
	imprimeX := closure()
	imprimeX()
}
