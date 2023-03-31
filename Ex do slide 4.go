package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Print("Digite uma string (frase):")
	fmt.Scan(&palavra)
	fmt.Println(strings.ToUpper(palavra))
}
