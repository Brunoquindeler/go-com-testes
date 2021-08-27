package iteracao

import "testing"

func TestRepetidor(t *testing.T) {
	repeticoes := Repeditor("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("\nEsperado: %v\nObtido: %v", esperado, repeticoes)
	}
}

func BenchmarkRepetidor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeditor("a")
	}
}
