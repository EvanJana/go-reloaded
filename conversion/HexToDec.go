package conv

import "strconv"

func HexToDec(words []string) ([]string, error) {
	res := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]
		last := len(res) - 1
		if len(res) > 0 && w == "(hex)" {
			nbr, err := strconv.ParseInt(res[last], 16, 64)
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