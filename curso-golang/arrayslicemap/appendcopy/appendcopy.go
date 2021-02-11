package main

import "fmt"

//funcoes append e copy usado no slice e array

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//o método append não é possível para array (já possui 3 posicoes definidas), pois o primeiro argumento precisa ser um slice (vazio)
	//exemplo: array1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6) //cria um novo "array" e adiciona os elementos
	fmt.Println(array1, slice1, len(slice1), len(array1)) 

	slice2 := make([]int, 2) //outro slice vazio, com duas posicoes
	copy(slice2, slice1) //copia para o slice2, os elementos do slice 1 / não é possível passar um array, somente slice para ambos
	fmt.Println(slice2) //Porém como o tamanho do 2 é menor que o 1, ele copia somente os dois primeiros elementos

}
