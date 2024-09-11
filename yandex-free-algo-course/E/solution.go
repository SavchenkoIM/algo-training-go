package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func scanMLText(r *bufio.Reader, lines int) string {
	sb := strings.Builder{}
	for i := 0; i < lines; i++ {
		line, _ := r.ReadString('\n')
		sb.WriteString(line)
	}
	return sb.String()
}

type RomanDigit struct {
	Role   int // 1 or 5
	Decade int // 0, 1, 2, 3...
}

var Vocabulary = map[rune]RomanDigit{
	'I': RomanDigit{1, 0},
	'V': RomanDigit{5, 0},
	'X': RomanDigit{1, 1},
	'L': RomanDigit{5, 1},
	'C': RomanDigit{1, 2},
	'D': RomanDigit{5, 2},
	'M': RomanDigit{1, 3},
}

func extractPrefix(number []rune) (int, []rune, int) { // number, rest, decade
	leading := number[0]
	v, ok := Vocabulary[leading]
	if !ok {
		return -1, number, 0
	}

	currDecade := v.Decade
	leadingLength := 1

	switch v.Role {
	case 1:
		for i := 1; i < len(number); i++ {
			if number[i] == leading {
				leadingLength++
				if leadingLength > 3 {
					return -1, number, 0
				}
			} else if Vocabulary[number[i]].Role == 5 && Vocabulary[number[i]].Decade == currDecade {
				if leadingLength > 1 {
					return -1, number, 0
				}
				return (5 - leadingLength) * int(math.Pow10(currDecade)), number[i+1:], currDecade
			} else if Vocabulary[number[i]].Role == 1 && Vocabulary[number[i]].Decade == currDecade+1 {
				if leadingLength > 1 {
					return -1, number, 0
				}
				return (10 - leadingLength) * int(math.Pow10(currDecade)), number[i+1:], currDecade
			} else {
				return leadingLength * int(math.Pow10(currDecade)), number[i:], currDecade
			}
		}
		return leadingLength * int(math.Pow10(currDecade)), nil, currDecade
	case 5:
		leadingLength = 0
		for i := 1; i < len(number); i++ {
			if Vocabulary[number[i]].Role == 1 && Vocabulary[number[i]].Decade == currDecade {
				leadingLength++
				if leadingLength > 3 {
					return -1, number, 0
				}
			} else {
				return (5 + leadingLength) * int(math.Pow10(currDecade)), number[i:], currDecade
			}
		}
		return (5 + leadingLength) * int(math.Pow10(currDecade)), nil, currDecade
	}
	return -1, number, 0
}

func convertRomanDigit(digit string) int {

	res := 0
	number := []rune(digit)

	n := 0
	d := 0
	prevDecade := 99999999
	sarr := number
	for len(sarr) > 0 {
		n, sarr, d = extractPrefix(sarr)
		res += n
		if prevDecade <= d {
			return -1
		}
		prevDecade = d
	}

	return res
}

func main() {

	in := bufio.NewReader(os.Stdin)

	dgt := strings.TrimSuffix(scanMLText(in, 1), "\n")

	fmt.Println(convertRomanDigit(dgt))

}
