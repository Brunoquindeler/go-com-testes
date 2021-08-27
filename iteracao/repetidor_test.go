package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetidor(t *testing.T) {
	repeticoes := Repetidor("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("\nEsperado: %v\nObtido: %v", esperado, repeticoes)
	}
}

func BenchmarkRepetidor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetidor("a", 5)
	}
}

func ExampleRepetidor() {
	r := Repetidor("a", 5)
	fmt.Println(r)
	// Output: aaaaa
}
