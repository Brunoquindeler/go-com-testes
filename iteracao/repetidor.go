package iteracao

// Repetidor recebe um caractere e a quantidade de vezes em que ele deve ser repetido
func Repetidor(caractere string, quantidadeRepeticoes int) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
