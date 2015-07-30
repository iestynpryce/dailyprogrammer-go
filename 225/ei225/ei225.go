package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename1 ... filenameN\n", os.Args[0])
	}

	for _, f := range os.Args[1:] {
		deColumize(f)
	}
}

func deColumize(filename string) {

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "[ERROR] file '%s' does not exist", filename)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var feature_plus bool = false
	var feature_pipe bool = false
	var firstline bool = true
	var linestart bool = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip the first line, info not needed
		if firstline {
			firstline = false
			continue
		}

		var buffer bytes.Buffer

		// Remove the features
		for _, c := range scanner.Text() {
			switch c {
			case '+':
				feature_plus = !feature_plus
				feature_pipe = false
				continue
			case '|':
				feature_pipe = !feature_pipe
				feature_plus = false
				continue
			}

			if !feature_plus && !feature_pipe {
				buffer.WriteRune(c)
			}
		}

		// Remove the colums

		stripped := strings.TrimSpace(buffer.String())
		line_length := len(stripped)

		if line_length == 0 {
			fmt.Println()
			linestart = true
		} else if !linestart {
			fmt.Print(" ")
		}

		for i, c := range stripped {
			switch c {
			case '-':
				if i+1 == line_length {
					linestart = true
					continue
				}
			}
			fmt.Printf("%c", c)
			linestart = false
		}
	}
}
