package conv

import (
	"strings"
)

func Ponctuation(words []string) []string {
	Apostrophes := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		quote := words[i]
		if quote == "'" || quote == `"` {
			matched := false
			for end := i + 2; end < len(words); end++ {
				if words[end] == quote {
					Apostrophes = append(Apostrophes, quote+strings.Join(words[i+1:end], " ")+quote)
					i = end
					matched = true
					break
				}
			}
			if matched {
				continue
			}
		}
		Apostrophes = append(Apostrophes, words[i])
	}

	text := strings.Join(Apostrophes, " ")
	res := ""
	punctuation := ".,!?:;"

	for i := 0; i < len(text); i++ {
		carac := text[i]
		if strings.ContainsRune(punctuation, rune(carac)) {
			res = strings.TrimRight(res, " ")
			res += string(carac)
			for i+1 < len(text) && text[i+1] == ' ' {
				i++
			}
			if i+1 < len(text) &&
				!strings.ContainsRune(punctuation, rune(text[i+1])) {
				res += " "
			}
		} else {
			res += string(carac)
		}
	}

	
	return strings.Fields(res)
}
