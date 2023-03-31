package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	var numero int
	fmt.Print("escreva uma palavra:")
	fmt.Scan(&palavra)
	fmt.Print("escreva um número:")
	fmt.Scan(&numero)

	transformado := strings.ToUpper(palavra[:numero]) + palavra[numero:]
	fmt.Println(transformado)
}
