package main

import "fmt"

//Não tem operador ternário / return nota >=6 ? "Aprovado" : "Reprovado" 
func obterResultado(nota float64) string /*retorna*/ {
	if nota >=6 {
	return "Aprovado"
	} 
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
