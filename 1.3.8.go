package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()
	var abbr []rune
	for _, word := range strings.Fields(phrase) {
		for _, rune := range word {
			if unicode.IsLetter(rune) {
				abbr = append(abbr, unicode.ToUpper(rune))
			}
			break
		}
	}
	fmt.Println(string(abbr))
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
