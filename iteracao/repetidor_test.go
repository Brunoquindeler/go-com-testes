package iteracao

import "testing"

func TestRepetidor(t *testing.T) {
	repeticoes := Repeditor("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Esperado: %v 	|	Obtido: %v", esperado, repeticoes)
	}
}
