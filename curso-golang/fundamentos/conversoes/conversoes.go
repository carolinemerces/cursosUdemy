package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(123)) //esse número correspondo ao da tabela 123 = {
	fmt.Println("Teste " + string(97))  // 97 = a

	//inteiro para string
	fmt.Println("Teste " + strconv.Itoa(123)) //int to string

	//string para int
	num, _ := strconv.Atoi("123") //string to int, retorno de dois valores (o segundo valor é um erro e para ignorá-la utiliza-se _)
	fmt.Println(num - 122)

	//conversão de string para boolean
	b, _ := strconv.ParseBool("true") //ou 1 
	if b {
		fmt.Println("Verdadeiro")
	}


}
