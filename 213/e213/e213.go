package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
	"A": "atta",
	"B": "bibbity",
	"C": "city",
	"D": "dickety",
	"E": "ebbity",
	"F": "fleventy",
}

func pronounce_hex(text string) {
	if strings.HasPrefix(text, "0x") { // check it's a hex value
		fmt.Print(text + " \"")
		for length, c := range text[2:] {
			switch {
			case length == 1 || length == 3 && c != '0':
				fmt.Print("-")
			case length == 2:
				fmt.Print("bitey ")
			}
			if c != '0' {
				trans := m[string(c)]
				if length == 1 || length == 3 {
					if !strings.Contains("onetwothreefourfivesixseveneightnine", trans) {
						trans = strings.ToLower(string(c)) + "ee"
					}
				}
				fmt.Print(trans)
				if length == 1 {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("\"")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pronounce_hex(strings.Trim(scanner.Text(), " "))
	}
}
