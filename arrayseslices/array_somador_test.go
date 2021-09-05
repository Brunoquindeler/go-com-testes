package arrayseslices

import (
	"fmt"
	"testing"
)

func TestArraySomador(t *testing.T) {
	numeros := [5]int{1, 2, 3, 4, 5}

	obtido := ArraySomador(numeros)
	esperado := 15

	if esperado != obtido {
		t.Errorf("\nDado: %v \nEsperado: %d \nObtido: %d", numeros, esperado, obtido)
	}
}

func ExampleArraySomador() {
	total := ArraySomador([5]int{1, 2, 3, 4, 5})
	fmt.Println(total)
	// Output: 15
}

func BenchmarkArraySomador(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArraySomador([5]int{1, 2, 3, 4, 5})
	}
}
