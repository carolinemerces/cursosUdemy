package main

import "fmt"

//funcao sem parâmetro e sem retorno
func f1() {
	fmt.Println("F1")
}

//funcao com parâmentro sem retorno
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s", p1, p2)
}

//funcao sem parâmetro, mas retorna algo
func f3() string {
	return "F3"
}

//funcao com 2 parâmetros e 1 retorno
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) //ao usar o return, para printar é necessário utilizar o fmt.Sprint
}

//funcao pura -> funcao que não mexe em nada externo à funcao / não influencia fora da funcao
//recebe um conjunto de paraâmetros e sempre gera (de forma determinística) o mesmo resultado de um conjunto de parâmetros

//uma funcao pode retornar multiplos valores (podendo ou não estar dentro de uma única estrutura)
//funcao sem parâmetro com 2 retornos
func f5() (string, string) {
	return "Retorno1", "Retorno 2"
}

//funcao main obrigatório para executar o programa
func main() {
	f1()
	f2("Parâmetro1", "Parâmetro 2\n")

	r3 := f3()
	fmt.Printf("F3: %s\n", r3)

	r4 := f4("Parâmetro1", "Parâmetro 2")
	fmt.Println(r4)

	r51, r52 := f5() //não é possível ignorar (sem o uso de _) um dos valores, se tratando de múltiplos valores, mas é possível ignorar totalmente o retorno de todos os valores
	fmt.Printf("F5: %s %s", r51, r52)
}

