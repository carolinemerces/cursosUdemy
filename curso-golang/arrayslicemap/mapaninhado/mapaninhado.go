package main

import "fmt"

//um map dentro de outro map - map aninhado

func main() {

	funcsPorLetra := map[string]map[string]float64{ //temos um map com chave string e como valor um outro map com chave string e valor float64

		"A": {
			"André Silva": 3457.78,
		}, //obrigatório uso de vírgula após as chaves também
		"J": {
			"João Junior": 4567.89,
		},
		"M": {
			"Maria Joaquina": 4990.90,
		},
	}

	delete(funcsPorLetra, "M")
	fmt.Println(funcsPorLetra)

	//varrer com for range, usando duas variáveis para atribuir a chave do primeiro map na letra e a chave o valor do segundo map no funcs
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	//for range dentro de outro for range - varrendo o primeiro map (chave) e o segundo map (valor dessa chave, que é outra chave e valor)
	for l := range funcsPorLetra {
		fmt.Printf("Letra: %s\n", l)

		for nome, salario := range funcsPorLetra [l] { //foi passado o valor da chave (nome e salario - segundo map), que era l (chave - primeiro map)
			fmt.Printf("%s salário: %.2f\n", nome, salario)
		}
	}

}
