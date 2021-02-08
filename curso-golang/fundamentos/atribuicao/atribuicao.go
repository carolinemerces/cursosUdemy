package main

import "fmt"

//atribuicao simples
func main() {
	var b byte = 3
	fmt.Println(b)

i := 3 //inferência
i += 3 // i = i + 3
i -=3 // i = i - 3
i /= 2 // i = i / 2
i *= 2 // i = i * 2
i %= 2 // i = i % 2

fmt.Println(i)

//atribuicão múltipla de valores
x, y := 1, 2 
fmt.Println(x, y)

//troca de valores de variáveis diretas
x, y = y, x
fmt.Println(x, y)
}


