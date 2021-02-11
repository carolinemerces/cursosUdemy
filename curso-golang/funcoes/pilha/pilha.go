package main

//pilha de execucao - stack
//uma funcao que chama a outra, que chama a outra, que chama a outra

import "runtime/debug"

//quarta funcao a ser chamada // terminando a execucao dessa funcao, ela volta para a funcao anterior que a qual foi chamada (f2 neste caso)
func f3(){
	debug.PrintStack() //imprime a pilha de excucao de um programa, no momento que ela for chamada
}

//terceira funcao a ser chamada //  terminando a execucao dessa funcao, ela volta para a funcao anterior que a qual foi chamada (f1 neste caso)
func f2(){
	f3()
}

//segunda funcao a ser chamada //  terminando a execucao dessa funcao, ela volta para a funcao anterior que a qual foi chamada (main neste caso)
func f1(){
	f2()
}

//funcao main inicia o programa //  terminando a execucao dessa funcao, a funcao termina e o programa conclui
func main(){
	f1()
}

//técnica recursão ou método recursivo (que chama ele mesmo várias vezes, estourando a pilha)