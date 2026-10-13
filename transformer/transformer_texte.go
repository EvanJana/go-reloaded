package text

import (
	"strconv"
	"strings"
	conv "github.com/EvanJana/go-reloaded/conversion"
)

func transform(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for _, w := range words {
		switch w {
		case "(low)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = conv.Lower(result[last])
			}
		case "(up)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = conv.Upper(result[last])
			}
		case "(cap)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = conv.Capitalize(result[last])
			}
		case "(hex)":
			if len(result) > 0 {
				last := len(result) - 1
				n, err := conv.HexToDec(result[last])
				if err == nil {
					result[last] = strconv.Itoa(n)
				}
			}
		case "(bin)":
			if len(result) > 0 {
				last := len(result) - 1
				n, err := conv.BinToDec(result[last])
				if err == nil {
					result[last] = strconv.Itoa(n)
				}
			}
		default:
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}
