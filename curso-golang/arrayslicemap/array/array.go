package main

import "fmt"

func main(){
	//estruturas estáticas (fixas, não variam de tamanho/posicoes), homogêneas (não variam de tipo, sempre o mesmo tipo de dado) e indexadas, comecando pelo zero

var notas [3]float64
fmt.Println(notas)

notas[0], notas[1], notas[2] = 7.8, 4.3, 5.0
fmt.Println(notas)

total := 0.0

for i := 0; i < len(notas); i++ {
	total += notas[i]
}

media := total / float64(len(notas))
fmt.Printf("Média %.2f\n", media)

}