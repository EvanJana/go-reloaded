package conv

import "strings"

func Capitalize(mot string) string {
	runes := []rune(strings.ToLower(mot))
	if len(runes) == 0 {
		return ""
	}
	return strings.ToUpper(string(runes[0])) + string(runes[1:])
}
