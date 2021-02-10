package main

import "fmt"

//switch multiplas selecoes, passando um valor que será analisado caso a caso

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota { //não há uma expressão que retorna verdadeiro ou falso, como no if ou for. Há uma selecao de setenca de código, mas a partir de um valor que será verificado em cada um dos casos
	case 10: //caso o valor inteiro seja 10, ele entra naquele bloco que está contido o valor e para a execucao, sem a necessidade da palavra break, mas sim fallthrough
		fallthrough // para continuar descendo para executar os próximos cases
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: 
		return "Nota inválida!"
	} 
}

func main () {
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(4))
}
