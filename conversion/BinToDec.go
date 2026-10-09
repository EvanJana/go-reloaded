package conv

import (
	"strconv"
)

func BinToDec(bin string) (int, error) {
	nbr, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(nbr), nil
}
