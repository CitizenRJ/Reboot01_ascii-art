package asciiArt

import (
	"fmt"
)

func PrintLines(input []int, str []string) []string {
	var word = make([]string, 8)
	for i := 0; i < len(input); i++ {
		input[i] = input[i] - 32
		// fmt.Println(input[i])
	}
	for i := 0; i < len(input); i++ {
		index := input[i]
		if index == 98 {
			word[7] = word[7] + "\n"
		} else if index < 0 || index >= len(str) {
			fmt.Printf("Invalid input code at index %d: %d\n", i, input[i])
			continue
		}
		for j := 0; j < 8; j++ {
			word[j] = word[j] + str[index*9+j+1]
		}

	}
	return word
}
