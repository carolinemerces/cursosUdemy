package main

import "fmt"

//método inspiração da oo, definição de uma função com possui um receiver e essa função ficar associada a uma estrutura de dados
//conceito de interface - se há uma estrutura que respeita os métodos de uma interface, sem deixar explícito, de forma implícita a linguagem entende que se trata de uma interface, essa estrutura tem métodos que sustentam os métodos de uma interface (polimorfismo)
//polimorfismo, a qual uma mesma variável se comportará de maneiras distintas dentro do programa, baseado no seu tipo interface, com método imprimível
//exemplo de dois métodos do tipo toString(interface)

type imprimivel interface {
	toString() string //para saber se uma estrutura é imprimível ou não, deve-se utilizar esse método associado (toString)
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

//interfaces são implementadas implicitamente
//função que recebe um reciever do tipo da struct acima
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %2.f", p.nome, p.preco)
}

//função que recebe como parâmetro uma interface com o método toString, ou seja, qualquer coisa que respeite essa interface, sem que haja uma relação direta
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Caroline", "das Mercês"}
	imprimir(coisa)

	coisa = produto{"Notebook Mac Pro", 1999.99} //polimorfismo (múltiplas formas) baseado em interface, uma hora a variável é pessoa e outra produto. Existe uma estrutura base (interface imprimível) de conceito (herança)
	fmt.Println(coisa.toString())

	//outro exemplo possível
	imprimir(produto{"Celular Sansung Galaxy", 999.99})

	//outro exemplo possível, com uma variável que não é imprimível
	p2 := produto{"Tablet Sansung Galaxy", 899.99}
	imprimir(p2)

}

