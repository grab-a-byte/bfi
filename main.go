package main

import (
	"fmt"
)

var bytes [30_000]byte

func main() {
	dc := 0
	ic := 0
	input := `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`
	inputLen := len(input)

	moveUntil := func(incChar, decChar rune, icIncrement int) {
		bracketCount := 1
		for bracketCount != 0 {
			ic += icIncrement
			if input[ic] == byte(incChar) {
				bracketCount++
			} else if input[ic] == byte(decChar) {
				bracketCount--
			}
		}
	}

	for ic = 0; ic < inputLen; ic++ {
		char := input[ic]
		switch char {
		case '<':
			if dc != 0 {
				dc--
			}
		case '>':
			if dc < 30_000 {
				dc++
			}
		case '+':
			bytes[dc]++
		case '-':
			bytes[dc]--
		case '.':
			fmt.Print(string(rune(bytes[dc])))
		case '[':
			{
				if bytes[dc] == 0 {
					moveUntil('[', ']', 1)
				}
			}
		case ']':
			{
				moveUntil(']', '[', -1)
				ic--
			}
		}
	}
}
