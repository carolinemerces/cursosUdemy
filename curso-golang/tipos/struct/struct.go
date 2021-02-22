package main

import "fmt"

//métodos - funcoes associadas a estruturas de dados
//um novo tipo (nome) e uma estrutura, dentro da estrutura coloca todos os dados
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//funcao ou método associado a estrutura acima
//receiver - receptor(dono) dessa funcao
func (p produto) precoComDesconto() float64 { //o produto é a struct a qual será o receptor / o p no ínicio é para referenciar o produto / e segue a funcao normalmente / nao é necessário parâmetros, pois estão dentro da struct
	return p.preco * (1 - (p.desconto)) //uso da notacao . para fazer o acesso do preco da struct, como o desconto
}

//funcao main, com declaracao da variável do tipo produto e atribuicao de valor para ele, com cálculo do desconto(funcao)
func main() {
	var produto1 produto
	produto1 = produto{ //para atribuicao de valor é necessário o uso de : e ,
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto()) //utilizacao do ponto para acessar a funcao

	var produto2 produto
	produto2 = produto{
		nome:     "Caneta",
		preco:    2.79,
		desconto: 0.03,
	}
	fmt.Println(produto2, produto2.precoComDesconto())

	//declaracao e atribuicao direta
	produto3 := produto{"Notebook", 2999.99, 0.5}
	fmt.Println(produto3, produto3.precoComDesconto())
}

