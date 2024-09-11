package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readString(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func readIntSlice(r *bufio.Reader) []int {
	res := make([]int, 0)
	for _, s := range strings.Fields(readString(r)) {
		i, _ := strconv.Atoi(s)
		res = append(res, i)
	}
	return res
}

type Range struct {
	Start int
	End   int
}

func main() {
	rdr := bufio.NewReader(os.Stdin)

	first := make([]Range, 0)
	second := make([]Range, 0)

	lFirst := readIntSlice(rdr)[0]
	for i := 0; i < lFirst; i++ {
		line := readIntSlice(rdr)
		first = append(first, Range{Start: line[0], End: line[1]})
	}

	lSecond := readIntSlice(rdr)[0]
	for i := 0; i < lSecond; i++ {
		line := readIntSlice(rdr)
		second = append(second, Range{Start: line[0], End: line[1]})
	}

	for len(first) > 0 && len(second) > 0 {
		rng := map[bool]*[]Range{true: &first, false: &second}
		_first := rng[first[0].Start <= second[0].Start]
		_second := rng[first[0].Start > second[0].Start]

		if (*_first)[0].End < (*_second)[0].Start {
			*_first = (*_first)[1:]
			continue
		}

		maxStart := max((*_first)[0].Start, (*_second)[0].Start)
		minEnd := min((*_first)[0].End, (*_second)[0].End)

		fmt.Println(maxStart, minEnd)

		if first[0].End <= minEnd {
			first = first[1:]
		}
		if second[0].End <= minEnd {
			second = second[1:]
		}
	}
}
