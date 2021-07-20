package main

import "fmt"

const (
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoOlaFrances   = "Bonjour, "
)

const (
	fr  = "fr"
	esp = "esp"
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case esp:
		prefixo = prefixoOlaEspanhol
	case fr:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("Bruno", ""))
}
