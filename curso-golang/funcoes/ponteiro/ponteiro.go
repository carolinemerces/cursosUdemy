package main

import "fmt"

//Em Go por padrão, passa-se cópia de valores ou uma referência passada como parâmetro

//primeira funcao receberá um valor inteiro
func inc1(n int) {
	n++
}

//segunda funcao receberá um ponteiro como parâmetro
//revisão -> um ponteiro é representado por *
func inc2(n *int) { //tipo ponteiro de inteiro
	//revisão: * neste caso é usado para desreferenciar, ou seja,
	//ter acesso ao valor no qual o ponteiro aponta/ referencia
	*n++ //vai retornar um valor e a partir dele incrementar, pois foi desreferenciado
	//atencao: não é possível incrementar um ponteiro
}

func main() {
	n := 1 
	inc1(n) //passagem por valor e não por referência, agindo em cima de uma cópia e não interferindo em n, nesse caso
	fmt.Println(n)

	//revisão: & usado para obter o endereco da variável
	inc2(&n) //passagem por referência/ o endereco dessa variavel, porém n sofrerá interferência, será incrememtado +1, devido ao desreferenciamento
	fmt.Println(n)
}

//atencao: evitar a passagem por referência e utilizar a passagem por valor