package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func readString(rdr *bufio.Reader) string {
	j, _ := rdr.ReadString('\n')
	return strings.Trim(j, "\r\n")
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	j := readString(rdr)
	s := readString(rdr)

	jr := []rune(j)
	sr := []rune(s)

	slices.Sort(jr)

	cnt := 0
	for _, s := range sr {
		_, f := slices.BinarySearch(jr, s)
		if f {
			cnt++
		}
	}
	fmt.Println(cnt)
}
