package main

import (
	"bufio"
	"fmt"
	"maps"
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

func main() {
	rdr := bufio.NewReader(os.Stdin)
	s1 := readString(rdr)
	s2 := readString(rdr)

	sm1 := map[rune]int{}
	for _, s := range []rune(s1) {
		sm1[s]++
	}
	sm2 := map[rune]int{}
	for _, s := range []rune(s2) {
		sm2[s]++
	}

	if maps.Equal(sm1, sm2) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
