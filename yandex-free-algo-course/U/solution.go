package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readString(rdr *bufio.Reader) string {
	j, _ := rdr.ReadString('\n')
	return strings.Trim(j, "\r\n")
}

func readIntSlice(rdr *bufio.Reader) []int {
	s := readString(rdr)
	res := make([]int, 0)
	for _, i := range strings.Fields(s) {
		ii, _ := strconv.Atoi(i)
		res = append(res, ii)
	}
	return res
}

var res = make([]string, 0)

func PrintBracketSequence(path string, opR, clR, opS, clS, ln int) {
	if len(path) == ln {
		res = append(res, path)
		return
	}

	// If open possible
	if opR+opS-clR-clS < ln-len(path) {
		// Pair of Square brackets cannot contain Round brackets
		if opS == clS {
			PrintBracketSequence(path+"(", opR+1, clR, opS, clS, ln)
		}
		PrintBracketSequence(path+"[", opR, clR, opS+1, clS, ln)
	}
	// If close Round possible
	if path[len(path)-1] != '[' && opR > clR && opS == clS {
		PrintBracketSequence(path+")", opR, clR+1, opS, clS, ln)
	}
	// If close Square possible
	if path[len(path)-1] != '(' && opS > clS {
		PrintBracketSequence(path+"]", opR, clR, opS, clS+1, ln)
	}
}

var lexicographOrder = map[string]int{"(": 2, "[": 3, ")": 4, "]": 5}

func stringToInt(s string) int {
	res := 0
	for i, c := range s {
		res += lexicographOrder[string(c)] * int(math.Pow10(len(s)-i-1))
	}
	return res
}

func sortFunc(s1, s2 string) int {
	i1 := stringToInt(s1)
	i2 := stringToInt(s2)
	if i1 > i2 {
		return 1
	}
	if i1 < i2 {
		return -1
	}
	return 0
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	n := readIntSlice(rdr)[0]
	if n%2 != 0 {
		return
	}

	PrintBracketSequence("(", 1, 0, 0, 0, n)
	PrintBracketSequence("[", 0, 0, 1, 0, n)

	slices.SortFunc(res, sortFunc)

	for _, s := range res {
		fmt.Println(s)
	}
}
