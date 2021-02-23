package main

import (
	"fmt"
	"strings"
)

//funcao, struct e receiver
//duas funcoes simulando os gets e sets
//funcao que lê o valor de uma estrutura e não altera ela
//Mas se quiser alterar é necessário um ponteiro(passar por referência/ponteiro)

type pessoa struct {
	nome string
	sobrenome string
}

//método nome completo
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//método alterar nome e sobrenome com a funcao Split
func (p *pessoa) setNomeCompleto(nomeCompleto string){ //receiver pessoa neste caso será um ponteiro para que seja possível sua alteracao / possui como parâmentro uma string e não retorna nada
	partes := strings.Split(nomeCompleto, " ") //a variável partes recebe o pacote strings e a funcao Split, recebendo como parâmetro a variável declarada na funcao setNomeCompleto, separando por espaco
	p.nome = partes[0] //o ponteiro vai receber a alteracao do nome completo, alterando somente o nome [0]
	p.sobrenome = partes[1]//o ponteiro vai receber a alteracao do nome completo, alterando somente o sobrenome [1]
}

func main(){
	p1 := pessoa{nome: "Caroline", sobrenome: "das Mercês"}
	fmt.Println(p1.getNomeCompleto())

	//a própria linguagem se encarrega de encapsular
	p1.setNomeCompleto("André Bispo") //para que a esses valores sejam de fato alterados é necessário o uso de um ponteiro na funcao set, caso contrário ela não será alterada com o novo valor adicionado / ele precisa passar por referência e não por valor
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Ana Júlia")
	fmt.Println(p1.getNomeCompleto())

	p2 := pessoa{nome: "Maria", sobrenome: "Joaquina"}
	fmt.Println(p2.getNomeCompleto())

	p2.setNomeCompleto("Ana Clara")
	fmt.Println(p2.getNomeCompleto())
}