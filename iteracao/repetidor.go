package iteracao

const quantidadeRepeticoes = 5

func Repeditor(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
