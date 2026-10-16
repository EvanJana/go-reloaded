package conv


func An(words []string) []string {
    res := []string{}

    for i := 0; i < len(words); i++ {
        w := words[i]

        if w == "(a)" && i+1 < len(words) && len(words[i+1]) > 0 {
            first := words[i+1]

            if first[0] == 'a' || first[0] == 'e' || first[0] == 'i' || first[0] == 'o' || first[0] == 'u' || first[0] == 'h' {
                res = append(res, "an")
            } else {
                res = append(res, w)
            }
        } else {
            res = append(res, w)
        }
    }

    return res
}
