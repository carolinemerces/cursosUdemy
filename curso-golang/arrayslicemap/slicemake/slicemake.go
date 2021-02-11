package main

import "fmt"

func main() {
	s := make([]int, 10) //funcao make - criacao de um slice com 10 elementos
	s[9] = 12 //elemento de índice 9 terá o valor 12
	fmt.Println(s)

	s = make([]int, 10, 20)//slice com 10 elementos, mas o array interno a ser criado terá 20 posicoes
	fmt.Println(s, len(s), cap(s)) //len tamanho / cap  capacidade máxima (tamanho do array interno) slice 10 / array 20

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8 , 9, 0) //adiciona mais 10 elementos atigindo a capacidade máxima/interna, mantendo o array interno como o mesmo
	fmt.Println(s, len(s), cap(s)) //slice 20 / array 20

	s = append(s, 1) //ao atingir a capacidade máxima, e adicionando mais um elemento, o slice recebe esse elemento e dobra a capacidade máxima interna do array 
	fmt.Println(s, len(s), cap(s)) //(slice 21 e array 40 neste caso)

}
