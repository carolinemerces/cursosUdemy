package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string { //i parâmetro genérico do tipo interface
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "Não faco ideia!!!"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(8))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
