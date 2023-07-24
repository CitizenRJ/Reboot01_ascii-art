package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func main() {
	Str := os.Args[1:]
	if len(Str) != 1 {
		if len(Str) == 0 {
			fmt.Println("Error: mising the word lol :).")
			os.Exit(0)
		} else if len(Str) > 2 {
			fmt.Println("Error: so many Args.")
			os.Exit(1)
		} else {
			fmt.Println("Error: mising output file name.")
			os.Exit(2)
		}
	} else {
		input := os.Args[1]
		file, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
		file.Close()
		output := os.Args[2]
		file, errs := os.Create(output)
		if errs != nil {
			fmt.Println("Failed to create file:", errs)
			os.Exit(2)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for i := 0; i < len(fileLines); i++ {
			if i < len(fileLines)-1 {
				fmt.Fprintln(writer, fileLines[i])
			} else {
				fmt.Fprint(writer, fileLines[i])
			}

		}
		writer.Flush()
		fmt.Println("Wrote to file: " + output + ".")
	}
}
