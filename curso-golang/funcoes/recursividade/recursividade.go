package main

import "fmt"

//recursividade = uma funcão que chama a si prórpia
//deve ter uma condicão de parada muito clara/definida para que ela não fique executando de forma indefinida e acabe estourando
//criacao de uma funcao com múltiplos valores e que retorna um erro

func fatorial(n int) (int, error) { //retorna um valor e um erro
	switch { //com valor verdadeiro
	case n < 0: //como não é possível ter um fatorial de número negativo, eu retorno uma mensagem de erro
		return -1, fmt.Errorf("Número inválido: %d", n) //retornou um número negativo para dizer que é inválido, além da mensagem de erro
	case n == 0: //o fatorial de 1 e 0, representam o resultado 1
		return 1, nil //não terá erro associado
	default: //nos demais casos, será chamado de forma recursiva a funcao, não possível retornar ela, mas é possível criar uma variável que receberá a funcão
		fatorialAnterior, _ := fatorial(n-1) //vai fazer uma chamada recursiva chamando a funcao fatorial com n-1, ignorando o erro
		return n * fatorialAnterior, nil //não terá erro associado
	}
}

func main() {
	//primeira chamada com um valor correto  - case n > 1
	resultado, _ := fatorial(5) //criar uma variável resultado, que recebe a funcão fatorial com o valor 5, ignorando o erro
	fmt.Println(resultado)

	//segunda chamada com um valor incorreto - case n < 0
	_, err := fatorial(-4) //ignora o valor do resultado, mas pega o valor do erro através da funcão fatorial que recebe como valor -4
	if err != nil { //tratando o erro
		fmt.Println(err)
	}

	//Uma solucão melhor seria...

}
