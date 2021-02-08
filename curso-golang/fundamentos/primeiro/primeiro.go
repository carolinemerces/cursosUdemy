//Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import "fmt"

/*A porta de entrada de um programa Go é a funcão main.
Essa funcao sempre será executada dentro do pacote main.
E por isso estão separados em pastas
*/
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
}

//Para executar -> control + alt(option) + N
