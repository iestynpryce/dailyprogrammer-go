package main

import (
	"fmt"
	"os"
	"strings"
)

// Returns true if the rune is an English vowel
func is_vowel(r rune) bool {
	return (r == 'A' ||
		r == 'a' ||
		r == 'O' ||
		r == 'o' ||
		r == 'E' ||
		r == 'e' ||
		r == 'I' ||
		r == 'i' ||
		r == 'U' ||
		r == 'u')
}

// Using the 'sub' string, substitute according to the rules of the game
func name_game_replacer(sub rune, s string) string {
	strlen := len(s)
	r := []rune(s)

	// If it starts with a single vowel, simply prepend the sub
	if strlen > 2 && is_vowel(r[0]) && !is_vowel(r[1]) {
		return string(sub) + strings.ToLower(s)
	}

	// If the sub and the first letter ar the same...
	if r[0] == sub {
		return string(r[0]) + "o-" + string(r[1:])
	} else {
		r[0] = sub
		return string(r)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Expected name on command line\n")
		os.Exit(1)
	}
	name := strings.Replace(os.Args[1], "!", "", 1)

	firstLineName := name_game_replacer('B', name)
	secondLineName := name_game_replacer('F', name)
	thirdLineName := name_game_replacer('M', name)

	fmt.Printf("%s, %s bo %s,\n", name, name, firstLineName)
	fmt.Printf("Bonana fanna fo %s,\n", secondLineName)
	fmt.Printf("Fee fy mo %s,\n", thirdLineName)
	fmt.Printf("%s!\n", name)
}
