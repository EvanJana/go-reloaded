package conv

func lower(s string) {
	rune := []rune(s)
	for _, caractere := range rune {
		if caractere >= 'A' && caractere <= 'Z' {
			caractere = caractere + 32
		}
	}
}
