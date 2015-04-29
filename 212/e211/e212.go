package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// define all constonants
var consonant []rune = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l',
	'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}

func contains(haystack []rune, needle rune) bool {
	for _, r := range consonant {
		if needle == r {
			return true
		}
	}
	return false
}

func encode(in []rune) (out []rune) {
	for _, c := range in {
		// creating a single char rune slice for use with unicode functions.
		out = append(out, c)
		if contains(consonant, unicode.ToLower(c)) {
			out = append(out, 'o', unicode.ToLower(c))
		}
	}
	return out
}

func decode(in []rune) (out []rune) {
	for i := 0; i < len(in); i++ {
		out = append(out, in[i])
		if contains(consonant, unicode.ToLower(in[i])) {
			i += 2
		}
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rs := encode([]rune(scanner.Text()))
		fmt.Println(string(rs))
		fmt.Println(string(decode(rs)))
	}
}
