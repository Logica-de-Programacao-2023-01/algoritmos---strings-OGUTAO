package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra, caractere string
	fmt.Print("Digite uma palavra:")
	fmt.Scan(&palavra)
	fmt.Print("Digite um caractere:")
	fmt.Scan(&caractere)

	count := strings.Count(palavra, caractere)
	fmt.Printf("O caractere '%s' aparece %d vezes na palavra '%s' .\n", caractere, count, palavra)
}
