package main

//funcao init -> por padrao é executado quando o arquivo em Go é interpretado

import "fmt"

func init(){ //essa funcao é inicializada antes da funcao main
	fmt.Println("Inicializando...")
}

func main(){
	fmt.Println("Main...")
}