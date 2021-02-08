package main

import (
	"fmt"
	m "math" //pode-se colocar a primeira letra do pacote a ser importado, para abreviar
	//ou ainda _ uma vez que esse pacote não seja utilizado na app, mas não queria retirá-lo
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida de criar var / declarando e inicializando a variável
	//a variável deve ser usada para não dar erro
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	//exemplos de blocos de construcao:
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//variável com atribuicao de valor
	var e, f bool = true, false
	fmt.Println(e, f)

	//exemplo de tipos diferentes de variáveis reduzidas
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
