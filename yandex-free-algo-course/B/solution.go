package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanInts(r *bufio.Reader) []int {
	line, _ := r.ReadString('\n')
	fields := strings.Fields(line)
	res := make([]int, len(fields))
	for i, f := range fields {
		rs, _ := strconv.Atoi(f)
		res[i] = rs
	}
	return res
}

func getMaxPoints(cards []int, mnMoves int) int {
	curr := 0
	for _, card := range cards[len(cards)-mnMoves:] {
		curr += card
	}
	res := curr

	for i := range mnMoves {
		curr += (cards[i] - cards[len(cards)-mnMoves+i])
		if curr > res {
			res = curr
		}
	}

	return res
}

func main() {

	in := bufio.NewReader(os.Stdin)

	scanInts(in)
	nMoves := scanInts(in)[0]
	cards := scanInts(in)

	fmt.Println(getMaxPoints(cards, nMoves))

}
