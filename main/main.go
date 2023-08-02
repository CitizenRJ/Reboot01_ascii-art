package main

import (
	"asciiArt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Str := os.Args[1:]
	if len(Str) > 1 || len(Str) < 1 {
		if len(Str) == 0 {
			fmt.Println("Error: no arguments :).")
			os.Exit(0)
		} else if len(Str) > 1 {
			fmt.Println("Error: so many Args.")
			os.Exit(1)
		}
	} else {
		var fileLines []string
		input := os.Args[1]
		if len(Str) == 1 {
			if len(Str) == 1 {
				file, err := os.Open("standard.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			}
		}
		var fileload []int
		result := strings.ReplaceAll(input, "\\t", "   ")
		result = strings.ReplaceAll(result, "\\n", " \\n ")
		words := strings.Split(result, " \\n ")
		for j := 0; j < len(words); j++ {
			if words[j] != "" {
				fileload = asciiArt.ConvertToASCIIArt(words[j])
				fileload1 := asciiArt.PrintLines(fileload, fileLines)
				for i := 0; i < len(fileload1); i++ {
					if strings.Join(fileload1[i], "") != "" {
						fmt.Println(strings.Join(fileload1[i], ""))
					}
				}
			}  else if j != 1 && words[j] == "" {
				fmt.Println()
			}
		}
	}
}
