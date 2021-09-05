package main

import (
	"fmt"

	"github.com/Brunoquindeler/go-com-testes/arrayseslices"
	"github.com/Brunoquindeler/go-com-testes/inteiros"
	"github.com/Brunoquindeler/go-com-testes/iteracao"
	"github.com/Brunoquindeler/go-com-testes/ola"
)

func main() {
	fmt.Println("--- Hello World! ---")

	fmt.Println("--------------------")

	fmt.Println("--- Olá ---")
	fmt.Println(ola.Ola("Bruno", ""))    // Olá, Bruno!!!
	fmt.Println(ola.Ola("Bruno", "fr"))  // Bonjour, Bruno!!!
	fmt.Println(ola.Ola("Bruno", "esp")) // Hola, Bruno!!!

	fmt.Println("--------------------")

	fmt.Println("--- Inteiros ---")
	fmt.Println("2 + 2:", inteiros.Adiciona(2, 2)) // 2 + 2: 4

	fmt.Println("--------------------")

	fmt.Println("--- Iteração ---")
	fmt.Println(iteracao.Repetidor("k", 5)) // kkkkk

	fmt.Println("--------------------")

	fmt.Println("--- Arrays e Slices ---")
	fmt.Println(arrayseslices.ArraySomador([5]int{1, 2, 3, 4, 5}))                        // 5
	fmt.Println(arrayseslices.SliceSomador([]int{2, 2, 4}))                               // 8
	fmt.Println(arrayseslices.MultiplesSlicesSomador([]int{2, 2, 4}, []int{1, 4}))        // [8 5]
	fmt.Println(arrayseslices.SomaTodoResto([]int{2, 2, 4}, []int{1, 4}, []int{3, 5, 4})) // [6 4 9]

}
