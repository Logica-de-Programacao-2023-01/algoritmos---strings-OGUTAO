package main

import "fmt"

func main() {
	var texto string
	fmt.Print("Escreva uma palavra:")
	fmt.Scan(&texto)
	s := texto
	fmt.Println("o comprimento da string é de:", len(s))
}
