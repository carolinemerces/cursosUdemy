package main

import "fmt"

//exemplo de interface esportiva com função turbo, com struct ferrari,  alterando um valor booleano

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { //é necessário utilizar um ponteir de ferrari para poder alterar o valor das variáveis
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} //é necessário atribuir o endereço de ferrari, devido ao uso da interface como tipo da variável
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
