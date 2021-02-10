package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //pacote randall (para gerar números aleatórios), pega a data atual e o nano segundo e passar como fonte. Sempre que o código for gerado, vai gerar um número novo 
	r := rand.New(s) 
	return r.Intn(10) //gera o número aleatório até 10
}

func main(){
	if i := numeroAleatorio(); i > 5 { //variável de inicializacao
		fmt.Println("Ganhou!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(i)
	}
}