package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{1, 2, 3} //array - tamanho fixo
	slice := []int{1, 2, 3}  //slice (uma parte) - tamanho variável / ponteiro que pega um elemento do array e a partir do tamanho desse slice ele determina uma estrutura contínua
	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	a1 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array, na verdade ele é um trecho/uma parte de um array
	s1 := a1[1:3] //vai do indíce 1 até o 2, pois não inclui o 3
	fmt.Println(a1, s1)
	s2 := a1[1:4]
	fmt.Println(a1, s2) //vai do indíce 1 até o 3, pois não inclui o 4

	s3 := a1[:2] //vai do início do indíce até o 1, pois não inclui o 2
	fmt.Println(a1, s3)

	s4 := s1[:1] //slice de um slice do mesmo array a1 - do zero ao 1, mas sem considerar o 1
	fmt.Println(a1, s4)
}
