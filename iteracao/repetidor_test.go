package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetidor(t *testing.T) {
	repeticoes := Repetidor("k", 5)
	esperado := "kkkkk"

	if repeticoes != esperado {
		t.Errorf("\nEsperado: %v\nObtido: %v", esperado, repeticoes)
	}
}

func BenchmarkRepetidor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetidor("k", 5)
	}
}

func ExampleRepetidor() {
	r := Repetidor("k", 5)
	fmt.Println(r)
	// Output: kkkkk
}
