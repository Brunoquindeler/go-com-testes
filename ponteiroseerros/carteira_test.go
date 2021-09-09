package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	obtido := carteira.Saldo()
	esperado := 10

	if esperado != obtido {
		t.Errorf("\nEsperado: %v \nObtido: %v", esperado, obtido)
	}

}
