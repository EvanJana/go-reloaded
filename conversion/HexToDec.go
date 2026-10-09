package conv

import "strconv"

func HexToDec(hex string) (int, error) {
	nbr, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(nbr), nil
}