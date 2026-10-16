package conv

import (
	"strconv"
	"strings"
)

func Capitalize(words []string) []string {
	res := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]
		last := len(res) - 1

		if len(res) > 0 && w == "(cap)" {
			if len(res[last]) > 0 {
				res[last] = strings.ToUpper(res[last][:1]) + strings.ToLower(res[last][1:])
			}

		} else if strings.HasPrefix(w, "(cap,") && len(res) > 0 {
			nStr := strings.TrimSuffix(strings.TrimPrefix(w, "(cap,"), ")")
			n, err := strconv.Atoi(nStr)

			if err != nil {
				res = append(res, w)
				continue
			}

			for j := 1; j <= n && j <= len(res); j++ {
				idx := len(res) - j
				if len(res[idx]) > 0 {
					res[idx] = strings.ToUpper(res[idx][:1]) + strings.ToLower(res[idx][1:])
				}
			}
		} else {
			res = append(res, w)
		}
	}

	return res
}
