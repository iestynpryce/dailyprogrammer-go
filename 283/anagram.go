package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	runner(os.Stdin)
}

// Having a runnier makes things a bit easier to test
func runner(reader *os.File) {
	// check that the file exists
	_, err := reader.Stat()
	if err != nil {
		panic(err)
	} else {
		buf := bufio.NewReader(reader)
		for {
			line, err := buf.ReadString('\n')
			if err != nil || line == "\n" {
				if err == io.EOF || len(line) > 0 {
					break
				}
				fmt.Fprintf(os.Stderr, "error reading line: %v", err)
				return
			}

			tokens := strings.Split(strings.Trim(line, "\n"), " ? ")
			if isanagram(tokens[0], tokens[1]) {
				fmt.Printf("%s is an anagram of %s\n", tokens[0], tokens[1])
			} else {
				fmt.Printf("%s is NOT an anagram of %s\n", tokens[0], tokens[1])
			}
		}
	}
}

// Using a map should make test O(n) assuming map access is O(1)
func isanagram(a, b string) bool {
	var m map[rune]int
	m = make(map[rune]int)

	astrip := strings.Map(compressAndLower, a)
	bstrip := strings.Map(compressAndLower, b)

	for _, c := range astrip {
		m[c]++
	}

	for _, c := range bstrip {
		if _, ok := m[c]; !ok {
			return false
		}
		m[c]--
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

// remove space and punctuation characters
func compressAndLower(r rune) rune {
	if unicode.IsSpace(r) || unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}
