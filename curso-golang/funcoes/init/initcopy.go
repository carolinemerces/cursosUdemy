package main

//a cópia de init.go é para mostrar que mesmo estando em outro arquivo, a segunda funcao init tbm se inicializa antes da funcao main que está no arquivo original

import "fmt"

func init(){ //essa funcao é inicializada depois da funcao init anterior e antes da funcao main
	fmt.Println("Inicializando Copy...")
}

//em uma mesma pasta, apenas um arquivo deve possuir a funcao main, para não causar conflito
//Lembrando que para executar ambos arquivos utilizar o terminal: go run *.go