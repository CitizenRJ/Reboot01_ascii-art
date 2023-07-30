package asciiArt

import (
	// "strings"
	"fmt"
)

func PrintLines(input []int, str []string) [][]string {
	// ha := len(input)
	var letters [][]string
	for i := 0; i < len(input); i++ {
		input[i] = input[i] - 32
		fmt.Println(input[i])
	}
	for i := 0; i < len(input); i++ {
		for j:= 0; j < 8; j++ {
			input[i] = input[i]*10
			letters[j][i] = append(letters[j], letters[i], str[input[i]+j])
		}
	}
	fmt.Println(letters)
	return letters
}
