package main

import "fmt"

//teste simples para provar que esse slice se trata de um mesmo array

func main() {

	s1 := make([]int, 10, 20) //slice 10 / array 20
	s2 := append(s1, 1, 2, 3) //slice 13 / array 20, com 1, 2, 3 adicionados ao final do slice
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2) //os dois slice apontam para o mesmo array interno


}
