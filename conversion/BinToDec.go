package conv

import "strconv"

func BinToDec(words []string) ([]string, error) {
	res := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]
		last := len(res) - 1
		if len(res) > 0 && w == "(bin)" {
			nbr, err := strconv.ParseInt(res[last], 2, 64)
			if err != nil {
				return nil, err
			}
			res[last] = strconv.Itoa(int(nbr))
		} else {
			res = append(res, w)
		}
	}

	return res, nil
}
