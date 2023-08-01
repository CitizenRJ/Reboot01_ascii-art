package main

import (
	"asciiArt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "regexp"
)

func main() {
	Str := os.Args[1:]
	if len(Str) > 2 || len(Str) < 1 {
		if len(Str) == 0 {
			fmt.Println("Error: no arguments :).")
			os.Exit(0)
		} else if len(Str) > 2 {
			fmt.Println("Error: so many Args.")
			os.Exit(1)
		}
		// else {
		// 	fmt.Println("Error: mising output file name.")
		// 	os.Exit(2)
		// }
	} else {
		var fileLines []string
		input := os.Args[1]
		if len(Str) == 1 {
			if len(Str) == 1 {
				// input := os.Args[1]
				file, err := os.Open("standard.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				// var fileLines []string
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			}
		} else {
			fileType := os.Args[2]
			shadow := "shadow"
			standard := "standard"
			thinkertoy := "thinkertoy"
			if fileType == standard {
				// input := os.Args[1]
				file, err := os.Open("standard.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				// var fileLines []string
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else if fileType == shadow {
				// input := os.Args[1]
				file, err := os.Open("shadow.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				// var fileLines []string
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else if fileType == thinkertoy {
				// input := os.Args[1]
				file, err := os.Open("thinkertoy.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				// var fileLines []string
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else {
				fmt.Println("Error: too many Args lol.")
				os.Exit(2)
			}
			// writer := bufio.NewWriter(file)
			// for i := 0; i < len(fileLines); i++ {
			// 	if i < len(fileLines)-1 {
			// 		fmt.Fprintln(writer, fileLines[i])
			// 	} else {
			// 		fmt.Fprint(writer, fileLines[i])
			// 	}
			// }
			// writer.Flush()
			// fmt.Println("Wrote to file: " + output + ".")
			// fmt.Println(fileLines)
		}
		var fileload []int
		result := asciiArt.ForNewLine(input)
		fmt.Println(result)
		words := strings.Split(result, " \\n ")
		// fmt.Println(len(words))
		// for _, word := range words {
			for j :=0 ; j < len(words) ; j++ {
			fileload = asciiArt.ConvertToASCIIArt(words[j])
			fileload1 := asciiArt.PrintLines(fileload, fileLines)
			// fmt.Println("\n \n")
			for i := 0; i < len(fileload1); i++ {
				// for j := 0; j < len(fileload1[i]); j++ {
					fmt.Printf(strings.Join(fileload1[i], "") + "\n")
				// }
				// fmt.Println()
			}
			// fmt.Println()
		}
		// fileload1 := asciiArt.PrintLines(fileload, fileLines)
		// fmt.Println("\n \n")
		// for i := 0; i < len(fileload1[0]); i++ {
		// 	for j := 0; j < len(fileload1); j++ {
		// 		fmt.Printf(strings.Join(fileload1[j], "") + "\n")
		// 	}
		// 	fmt.Println()
		// }
		// for i := 0; i < len(fileload1); i++ {
		// 	// fmt.Printf(fileload1[i] + "$\n")
		// 	fmt.Printf(strings.Join(fileload1[i], "") + "$\n")
		// }
		// fmt.Println(fileLines[11])
	}
}
