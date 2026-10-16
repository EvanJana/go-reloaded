package main

import (
	"os"
	"strings"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lignes := strings.Split(string(data), "\n")
	for i, ligne := range lignes {
		words := strings.Fields(ligne)
		words, err = conv.BinToDec(words)
		if err != nil {
			return
		}
		words, err = conv.HexToDec(words)
		if err != nil {
			return
		}

		words = conv.Upper(words)
		words = conv.Lower(words)
		words = conv.Capitalize(words)
		words = conv.An(words)
		words = conv.Ponctuation(words)
		lignes[i] = strings.Join(words, " ")
	}

	resultat := strings.Join(lignes, "\n")
	os.WriteFile("result.txt", []byte(resultat), 0644)
}
