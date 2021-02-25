package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Arquivo corrompido!")
	if err != nil {
		fmt.Print(err)
	}
}