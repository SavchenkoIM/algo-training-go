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

func main() {
	rdr := bufio.NewReader(os.Stdin)
	n := readIntSlice(rdr)[0]
	last := 0
	for i := 0; i < n; i++ {
		d := readIntSlice(rdr)[0]
		if i == 0 || d != last {
			fmt.Println(d)
		}
		last = d
	}
}
