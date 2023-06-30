package main

import "fmt"

func main() {
	var s1 string
	var s2 string
	fmt.Print("Digite o primeiro string:")
	fmt.Scan(&s1)
	fmt.Print("Digite o segundo string:")
	fmt.Scan(&s2)

	s3 := s1 + " " + s2
	fmt.Println(s3)

}
