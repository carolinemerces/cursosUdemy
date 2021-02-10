package main

import (
	"fmt"
	"time"
)

//switch sem nenhuma condicao associada a ele

func main() {
	t := time.Now() //hora atual
	switch true {   //switch true e ele vai procurar o primeiro case que seja verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
