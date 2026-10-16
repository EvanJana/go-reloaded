package conv

import (
    "strings"
    "strconv"
)

func Upper(words []string) []string {
    res := []string{}

    for i := 0; i < len(words); i++ {
        w := words[i]

        if strings.HasPrefix(w, "(up,") && len(res) > 0 {
            nStr := strings.TrimSuffix(strings.TrimPrefix(w, "(up,"), ")")
            n, err := strconv.Atoi(nStr)
            if err != nil {
                res = append(res, w)
                continue
            }

            for j := 1; j <= n && j <= len(res); j++ {
                res[len(res)-j] = strings.ToUpper(res[len(res)-j])
            }
        } else if len(res) > 0 && w == "(up)" {
            res[len(res)-1] = strings.ToUpper(res[len(res)-1])
        } else {
            res = append(res, w)
        }
    }

    return res
}