package asciiArt

import (
	// "strings"
	// "fmt"
)

func ConvertToASCIIArt(input string) []int {
	// Define the ASCII characters for each character in the input string
	// You can customize this based on your desired representation
	// asciiChars := map[byte]string{
	// 	'H': "@",
	// 	'e': "B",
	// 	'l': "$",
	// 	'o': "&",
	// 	',': "^",
	// 	' ': " ",
	// 	// Add more mappings for other characters
	// }
	// Split the input string into lines
	// lines := strings.Split(input, "\n")
	// Convert each line to ASCII art
	var word []int
	// for _, line := range lines {
	// Convert the characters in the line to ASCII art
	for i := 0; i < len(input); i++ {
		// fmt.Println(i)
		char := int(rune(input[i]))
		// fmt.Println(char)
		if input[i] == '\\' && i < len(input) {
			if input[i+1] == 'n' {
				char = 127
				word = append(word, char)
				i = i + 1
				// fmt.Println(i)
			} else {
				word = append(word, char)
			}
		} else {
			word = append(word, char)
		}
		// asciiChar, exists := asciiChars[char]
		// if exists {
		// 	output += asciiChar
		// } else {
		// 	// Handle unknown characters
		// 	output += "?"
		// }
	}
	// Add a line break after each line
	// output += "\n"
	// }
	// fmt.Println(word)
	return word

}
