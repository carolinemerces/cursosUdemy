package main

import "fmt"

//funcao variatica -> receberá uma quantidade variada de números e fará a média desses (não importando quantos)
func media(numeros ...float64) float64 { //representado pelos ..., pois pode receber variados parâmetros dentro de uma funcao
	total := 0.0                  //esses parâmetros serão passados como um array
	for _, num := range numeros { //utilizacao do for range, ignorando o índice e pegando os números dentro do parâmetro e colocando na variável num através do range
		total += num //a variável total receberá o número e fará a soma do próximo número
	}
	return total / float64(len(numeros)) //aqui foi necessário converter o tamanho do array em float64, devido os tipos de dados a serem trabalhados tbm serem float64. Operacão necessária para conseguir fazer a média, pegando o tamanho do array (len) / inteiro -> float64
}

func main() {
	fmt.Printf("Média: %.2f", media(7.0, 4.5, 5.8, 8.9))
}
