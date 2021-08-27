package iteracao

func Repeditor(caractere string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes = repeticoes + caractere
	}

	return repeticoes
}
