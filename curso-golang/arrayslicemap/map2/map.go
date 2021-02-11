package main

//Outro exemplo de map, como atribuir chave e valor

import "fmt"

func main() {

	funcsESalarios := map[string]float64 {
		"José João":   1132.00, //uso da vírgula é obrigatório
		"Maria Maria": 3232.00,
		"Ana Clara":   2032.00,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["Rael Fofi"] = 5000.60 //fora do bloco acima, como já foi inicializado, não necessita de vírgula
	fmt.Println(funcsESalarios)

	//tentar excluir algum elemento que não existe, não acontecerá nada
	delete(funcsESalarios, "Inexistente")
	fmt.Println(funcsESalarios)

	//varrer usando variáveis para especificar chave e valor
	for nome, salario := range funcsESalarios {
		fmt.Printf("%s salário: %.2f\n", nome, salario)
	}

}
