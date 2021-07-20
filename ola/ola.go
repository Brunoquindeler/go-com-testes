package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

const frances = "frances"
const espanhol = "espanhol"

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == frances {
		return prefixoOlaFrances + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
