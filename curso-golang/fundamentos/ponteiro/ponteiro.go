package main

import "fmt"

func main() {
	i := 1

	//Go não tem operacoes aritméticoas com ponteiro
	//O ponteiro é uma referência de memória (um endereco) para várias variáveis
	//Em Go é feito através do * ao lado da variável que você quer acessar o endereco como ponteiro
	//programacao orientada a lugar (vídeo no youtube)

	var p *int = nil //int - tipo do ponteiro (início do dado)
	p = &i           //pegue o endereco da variável i e atribua ao p (ponteiro)

	*p++ //pegou o valor associado ao ponteiro (desreferenciando) e incrementando +1
	i++

	fmt.Println(p, *p, i, &i)
}
