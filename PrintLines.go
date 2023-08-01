package asciiArt

import (
	"fmt"
)

// func PrintLines(input []int, str []string) []string {
// 	var word = make([]string, 0)
// 	for i := 0; i < len(input); i++ {
// 		input[i] = input[i] - 32
// 		// fmt.Println(input[i])
// 	}
// 	for i := 0; i < len(input); i++ {
// 		index := input[i]
// 		if index == 95 {
// 			word = append(word, "\n")
// 		} else if index < 0 || index >= len(str) {
// 			fmt.Printf("Invalid input code at index %d: %d\n", i, input[i])
// 			continue
// 		}
// 		for j := 0; j < 8; j++ {
// 			if index == 95 {
// 				word = append(word, "\n")
// 				break
// 			} else {
// 				word = append(word, str[index*9+j+1])
// 			}
// 		}
// 	}
// 	return word
// }
func PrintLines(input []int, str []string) [][]string {
	word := make([][]string, 8)
	for i := 0; i < len(input); i++ {
		input[i] = input[i] - 32
	}
	for i := 0; i < len(input); i++ {
		index := input[i]
		if index == 95 {
			word[7] = append(word[7], "")
		} else if index < 0 || index >= len(str) {
			fmt.Printf("Invalid input code at index %d: %d\n", i, input[i])
			continue
		}
		for j := 0; j < 8; j++ {
			if index == 95 {
				word[7] = append(word[7], "\n")
				break
			} else {
				word[j] = append(word[j], str[index*9+j+1])
			}
		}
	}
	return word
}
