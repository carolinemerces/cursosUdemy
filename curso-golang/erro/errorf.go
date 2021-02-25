package main

import (
	"fmt"
)

func main() {
	const name, id = "Caroline", 1
	err := fmt.Errorf("usuário %q (id %d) não encontrado", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
