package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Bruno")
	esperado := "Olá, Bruno"

	if resultado != esperado {
		t.Errorf("Resultado: '%s' | Esperado: '%s'", resultado, esperado)
	}
}
