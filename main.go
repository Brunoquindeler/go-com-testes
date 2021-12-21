package main

import (
	"fmt"

	"github.com/Brunoquindeler/go-com-testes/arrayseslices"
	"github.com/Brunoquindeler/go-com-testes/estruturasmetodosinterfaces"
	"github.com/Brunoquindeler/go-com-testes/inteiros"
	"github.com/Brunoquindeler/go-com-testes/iteracao"
	"github.com/Brunoquindeler/go-com-testes/ola"
	"github.com/Brunoquindeler/go-com-testes/ponteiroseerros"
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

	fmt.Println("--------------------")

	fmt.Println("--- Estruturas, Métodos e Interfaces ---")
	retangulo := estruturasmetodosinterfaces.Retangulo{Largura: 12, Altura: 6}
	fmt.Println(retangulo.Area()) // 72
	circulo := estruturasmetodosinterfaces.Circulo{Raio: 10}
	fmt.Println(circulo.Area()) // 314.1592653589793
	triangulo := estruturasmetodosinterfaces.Triangulo{Base: 12, Altura: 6}
	fmt.Println(triangulo.Area()) // 36

	fmt.Println("--------------------")

	fmt.Println("--- Ponteiros e Erros ---")
	carteira := ponteiroseerros.Carteira{}
	fmt.Println(carteira.Saldo().String())          // 0 BTC
	carteira.Depositar(ponteiroseerros.Bitcoin(10)) // Adiciona 10 ao saldo.
	fmt.Println(carteira.Saldo().String())          // 10 BTC
	carteira.Sacar(ponteiroseerros.Bitcoin(5))      // Retira 5 do saldo.
	fmt.Println(carteira.Saldo().String())          // 5 BTC
	carteira.Sacar(ponteiroseerros.Bitcoin(50))     // Não é possível realizar essa transação e o saldo permanece o mesmo pois é insuficiente.
	fmt.Println(carteira.Saldo().String())          // 5 BTC
}
