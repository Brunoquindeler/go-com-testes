package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%s' | Esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Bruno", "")
		esperado := "Olá, Bruno"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Diz 'Olá, Mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
