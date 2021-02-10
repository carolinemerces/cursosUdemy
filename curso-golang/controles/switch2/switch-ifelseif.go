package main

import "fmt"

func main() {
	n := 8.6
	switch {
	case n >= 9 && n <= 10:
		//fmt.Printf("Nota: %d - A", n)
		fmt.Println("A")
	case n >= 8 && n <= 9:
		//fmt.Printf("Nota: %d - B", n)
		fmt.Println("B")
	case n >= 5 && n <= 8:
		//fmt.Printf("Nota: %d - C", n)
		fmt.Println("C")
	case n >= 3 && n <= 5:
		//fmt.Printf("Nota: %d - D", n)
		fmt.Println("D")
	default:
		//fmt.Printf("Nota: %d - E", n)
		fmt.Println("E")
	}
}
