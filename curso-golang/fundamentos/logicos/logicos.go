package main

import "fmt"

//funcoes múltiplos valores (os operadores abaixo operam sobre duas variáveis, portanto são binários [trab1 e trab2])
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete /*negacao é unário, pois opera em cima de uma variável*/)
}
