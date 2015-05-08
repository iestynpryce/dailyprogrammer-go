package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var first_byte = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
	"A": "a",
	"B": "bee",
	"C": "cee",
	"D": "dee",
	"E": "ee",
	"F": "eff",
}

var second_byte = map[string]string{
	"10": "ten",
	"11": "eleven",
	"12": "twelve",
	"13": "thirteen",
	"14": "fourteen",
	"15": "fifteen",
	"16": "sixteen",
	"17": "seventeen",
	"18": "eighteen",
	"19": "nineteen",
	"20": "twenty",
	"30": "thirty",
	"40": "fourty",
	"50": "fifty",
	"60": "sixty",
	"70": "seventy",
	"80": "eighty",
	"90": "ninety",
	"A":  "atta",
	"B":  "bibbity",
	"C":  "city",
	"D":  "dickety",
	"E":  "ebbity",
	"F":  "fleventy",
}

func pronounce_twochar(text string, i *int, d int) {
	if strings.Contains("123456789", string(text[*i])) && d > 1 {
		if text[*i] == '1' &&
			strings.Contains("0123456789", string(text[*i+1])) {
			fmt.Print(second_byte[text[*i:*i+2]])
			*i++
			if d > 3 {
				fmt.Print(" bitey ")
			}
		} else {
			fmt.Print(second_byte[string(text[*i])+"0"])
			next := string(text[*i+1])
			if next != "0" || d > 3 {
				fmt.Print("-")
			}
		}
	} else {
		fmt.Print(second_byte[string(text[*i])])
		if string(text[*i]) != "0" && string(text[*i+1]) != "0" {
			fmt.Print("-")
			fmt.Print(first_byte[string(text[*i+1])])
			if d > 3 {
				fmt.Print(" bitey ")
			}
			*i++
		} else if d > 3 {
			fmt.Print("-")
		}
	}
}

func pronounce_hex(text string) {
	if strings.HasPrefix(text, "0x") { // check it's a hex value
		fmt.Print(text + ": ")

		subtext := text[2:]
		nchars := len(subtext)

		var last_zero bool = false

		if subtext[0] == '0' {
			last_zero = true
		}

		for i := 0; i < nchars; i++ {
			d := nchars - i

			if last_zero && subtext[i] == '0' {
				continue
			} else {
				last_zero = false
			}

			switch d {
			case 4:
				{
					pronounce_twochar(subtext, &i, d)
				}
			case 3:
				{
					s := string(subtext[i])
					fmt.Print(first_byte[s])
					if s != "0" {
						fmt.Print(" ")
					}
					fmt.Print("bitey ")
				}
			case 2:
				{
					pronounce_twochar(subtext, &i, d)
				}
			case 1:
				{
					s := string(subtext[i])
					if s != "0" {
						fmt.Print(first_byte[s])
					}
				}
			}
		}
		if last_zero {
			fmt.Print("zero")
		}
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pronounce_hex(strings.Trim(scanner.Text(), " "))
	}
}
