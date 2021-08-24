package main

import (
	"fmt"

	"github.com/Brunoquindeler/go-com-testes/inteiros"
	"github.com/Brunoquindeler/go-com-testes/ola"
)

func main() {
	fmt.Println("--- Hello World! ---")

	fmt.Println("--------------------")

	fmt.Println("--- Olá ---")
	fmt.Println(ola.Ola("Bruno", ""))
	fmt.Println(ola.Ola("Bruno", "fr"))
	fmt.Println(ola.Ola("Bruno", "esp"))

	fmt.Println("--------------------")

	fmt.Println("--- Inteiros ---")
	fmt.Println("2 + 2:", inteiros.Adiciona(2, 2))

}
