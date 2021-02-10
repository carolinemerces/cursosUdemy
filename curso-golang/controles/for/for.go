package main

import (
	"fmt"
	"time"
)

func main () {

	//for similiar ao while
	i := 1 // contador
	for i<= 10 { //não tem while em Go, mesmo em condicoes indeterminadas de repeticoes se usa o for
		fmt.Println(i)
		i++
	}

	//for tradicional
	for i := 0; i <=20; i += 2 {
		//fmt.Println(i)
		fmt.Printf("%d ", i)
	}

	//for tradicional com um if dentro de um for
	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}

	//laco infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) //para imprimir esperando 1 segundo
		// time.Sleep(time.Second * 5) para imprimir esperando 5 segundo
	}
}

//Veremos o foreach no capítulo de array