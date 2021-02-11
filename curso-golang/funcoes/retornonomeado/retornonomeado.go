package main

import "fmt"

//retorno nomeado -> atribuir nomes para os retornos
//método de troca de variáveis - uma pela outra
func trocar(p1, p2 int) (segundo, primeiro int) /*retorno nomeado*/ { //como o retorno é nomeado, ele vai retornar da forma que foi colocada dentro dos parênteses, não importando a ordem das atribuicões embaixo
	segundo = p2
	primeiro = p1
	return //retorno limpo, quando eu já atribui o que quero retornar, não precisamos colocar novamente o que retornar
}

func main() {

	x, y := trocar(2, 3)
	fmt.Println(x, y)

	//não há problema em ter variáveis com o mesmo nome do retorno nomeado, desde que declaradas apenas uma vez dentro do programa
	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
