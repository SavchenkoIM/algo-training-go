package main

import (
	"bufio"
	"fmt"
	"os"
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

func printBracketsSet(maxBrackets int, currOpen int, currClose int, path []byte) {
	if currOpen == maxBrackets {
		for i := currClose; i < maxBrackets; i++ {
			path = append(path, byte(')'))
		}
		fmt.Println(string(path))
		return
	}
	printBracketsSet(maxBrackets, currOpen+1, currClose, append(path, '('))
	if currOpen > currClose {
		printBracketsSet(maxBrackets, currOpen, currClose+1, append(path, ')'))
	}
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	n := readIntSlice(rdr)[0]
	printBracketsSet(n, 0, 0, []byte{})
}
