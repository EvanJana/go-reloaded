package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	os.WriteFile("result.txt", data, 0644)
	result, err := os.ReadFile("result.txt")
	if err != nil {
		return 
	}
	os.Stdout.Write(result)

	/*lignes := strings.Split(string(result), "\n")
	for _, ligne := range lignes {
		mots := strings.Fields(ligne)
	}*/
}