package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// define all constonants
var consonant []byte = []byte{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l',
	'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}

func encode(in []byte) (out []byte) {
	for _, c := range in {
		// creating a single char byte slice for use with bytes functions.
		buf := []byte(string(c))
		out = append(out, c)
		if bytes.Contains(consonant, bytes.ToLower(buf)) {
			out = append(out, 'o', bytes.ToLower(buf)[0])
		}
	}
	return out
}

func decode(in []byte) (out []byte) {
	for i := 0; i < len(in); i++ {
		buf := []byte(string(in[i]))
		out = append(out, in[i])
		if bytes.Contains(consonant, bytes.ToLower(buf)) {
			i += 2
		}
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ba := encode(scanner.Bytes())
		fmt.Println(string(ba))
		fmt.Println(string(decode(ba)))
	}
}
