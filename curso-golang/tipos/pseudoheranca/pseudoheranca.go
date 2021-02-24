package main

import "fmt"

//não existe heranca em estruturas, mas pseudo-heranca
//dentro de uma estrutura utilizar atributos de outras e parecer que está herdando -> composicao
//colocar um carro dentro de uma categoria

//identificador e um tipo
type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro            //campo anonimo, sem identificador, mas do tipo carro(struct)
	turboLigado bool //identificador e tipo
}

func main() {
	f := ferrari{} //dentro de ferrari não há o atributo nome, mas sim no carro, e pelo fato de ter criado um atributo anonimo é possível acessar diretamente o atributo nome
	f.nome = "F40"
	f.velocidade = 0
	f.turboLigado = true

	fmt.Printf("A Ferrari %s está com o turbo ligado? %t ", f.nome, f.turboLigado)//é possível ler e alterar esses dados
	fmt.Println(f) //composicao de uma estrutira maior (true), com identificador, e com uma estrutura interna (carro), sem identificador com campo anonimo
}

//Foi possível alterar os atributos recebidos de carro, não por heranca, mas por composicao, usando o campo anonimo(utilizando o tipo)