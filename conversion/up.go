package conv

import (
	"strconv"
)

func up(s string){
	rune := []rune(s)
	for _, caractere := range rune {
		if caractere >= 'a' && caractere <= 'z' {
			caractere = caractere - 32
		}
	}
}