package main

//map - chave - valor -> uma chave que não repeticao, dois elementos com a mesma chave, o elemento valor vai ser substituído pelo último adicionado, ambos possuem tipos homogêneos

import "fmt"

func main(){

		//var aprovados map[int]string //a chave vai inteiro e o valor uma string
		//maps devem ser inicializados

		aprovados := make(map[int]string) //inicializado através da funcão make com uma estrutura map

		//estrutura muito próxima do array, mas ao invés de usar um índice, passa-se uma chave, podendo de ser do tipo que desejar
		aprovados[12345678909] /*chave*/ = "Maria" /*valor*/
		aprovados[12345678888] = "Carol"
		aprovados[12345678000] = "Ana"

		fmt.Println(aprovados)

		for cpf /*chave*/, nome /*valor*/ := range aprovados { //for range -> para percorrer a chave e o valor de map, atribuindo esses valores as variáveis nome e cpf, podendo ignorar uma das duas usando o _
			fmt.Printf("%s (CPF: %d)\n", nome, cpf) 
		}

		//ver o valor de uma chave
		fmt.Println(aprovados[12345678000])

		//excluir o valor de uma chave
		delete(aprovados, 12345678000) //parâmetros (map, chave)
		fmt.Println(aprovados[12345678000])
		


}