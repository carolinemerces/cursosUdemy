package main

import "fmt"

//struct que possui outra struct como parte de sua estrutura, ou seja struct aninhada

//primeira struct item
type item struct {
	produtoID int
	qtde      int
	preco     float64
}

//segunda struct produto
type pedido struct {
	userID int
	itens  []item //dentro da struct de pedido tem a struct de item, e como serão vários pedido, colocou-se um slice(tamanho não defenido) de item
}

//método: funcao valorTotal com a segunda struct como referido
func (p pedido) valorTotal() float64 { //receiver p pedido / sem parâmetros, pois estão dentro das structs
	total := 0.0
	for _, item := range p.itens { //percorrer o slice de item (p.itens) através do for range / ignorando o índice
		total += item.preco * float64(item.qtde) //+= contribuicao aditiva / a qtde de itens é inteiro e por isso deve ser convertido para o cálculo com o float 64
	}
	return total
}

func main() { //atribuicao de valores na variável do tipo pedido com atributos de ambas structs
	pedido := pedido{
		userID: 1,
		itens: []item{ //slice de item (vários itens)
			item{produtoID: 1, qtde: 2, preco: 12.10}, 
			item{2, 1, 23.49},//produtoID, qtde, preco
			item{11, 200, 3.34}, //produtoID, qtde, preco
		},
	}
	fmt.Printf("Valor total do pedido é R$ %2.f", pedido.valorTotal()) //método que recebe um receiver do tipo pedido, percorreu os itens de uma outra struct, fez o cálculo e retornou o valor total

}
