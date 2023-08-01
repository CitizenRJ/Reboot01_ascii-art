package asciiArt

import (
	// "fmt"
	"regexp"
)

func ForNewLine(Str string) string {
	rege := regexp.MustCompile(`\\s*?(\n)\\s*?`)
	result := rege.ReplaceAll([]byte(Str), []byte(" $1 "))
	return string(result)
}
