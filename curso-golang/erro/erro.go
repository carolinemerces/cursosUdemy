package main

import (
	"fmt"
	"time"
)
//Exemplo de tratamento de erro na documentacao de Go: golang.org/pkg/errors/
// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(2021, 2, 25, 17, 00, 0, 0, time.UTC),
		"Tratamento de erro",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
