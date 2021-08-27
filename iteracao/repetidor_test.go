package iteracao

import "testing"

func TestRepetidor(t *testing.T) {
	repeticoes := Repeditor("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("\nEsperado: %v\nObtido: %v", esperado, repeticoes)
	}
}
