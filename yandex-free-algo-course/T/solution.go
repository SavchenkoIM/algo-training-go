package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readString(buf *bufio.Reader) string {
	str, _ := buf.ReadString('\n')
	return strings.Trim(str, "\r\n")
}

func readIntSlice(buf *bufio.Reader) []int {
	str := readString(buf)
	fields := strings.Fields(str)
	res := make([]int, 0, len(fields))
	for _, f := range fields {
		i, _ := strconv.Atoi(f)
		res = append(res, i)
	}
	return res
}

type BinaryRepresentation struct {
	val int
}

func (b BinaryRepresentation) Len() int {
	res := 0
	for b.val > 0 {
		b.val = b.val >> 1
		res++
	}
	return res
}

func (b BinaryRepresentation) BitAt(i int) bool {
	return b.val>>i&1 == 1
}

func cmpBinary(a, b BinaryRepresentation) int {
	if a.val > b.val {
		return 1
	}
	if a.val < b.val {
		return -1
	}
	return 0
}

func cmpBinaryByPos(pos int) func(a, b BinaryRepresentation) int {
	return func(a, b BinaryRepresentation) int {
		if a.BitAt(pos) && !b.BitAt(pos) {
			return 1
		}
		if !a.BitAt(pos) && b.BitAt(pos) {
			return -1
		}
		return 0
	}
}

func bestMatch(val BinaryRepresentation, pos int, subVals []BinaryRepresentation) int {
	sum := 0
	for pos > 0 {
		onesStart, _ := slices.BinarySearchFunc(subVals,
			BinaryRepresentation{val: 1 << (pos - 1)},
			cmpBinaryByPos(pos-1))
		ones := subVals[onesStart:]
		zeros := subVals[:onesStart]
		if len(zeros) == 0 && len(ones) == 0 {
			break
		}
		if val.BitAt(pos - 1) {
			if len(zeros) > 0 {
				sum += 1 << (pos - 1)
				subVals = zeros
			} else {
				subVals = ones
			}
		} else {
			if len(ones) > 0 {
				sum += 1 << (pos - 1)
				subVals = ones
			} else {
				subVals = zeros
			}
		}

		pos--
	}
	return sum
}

func maxXor(vals []int) int {
	slices.Sort(vals)

	bVals := make([]BinaryRepresentation, 0)
	for _, val := range vals {
		bVals = append(bVals, BinaryRepresentation{
			val: val,
		})
	}

	maxLen := 0
	for _, val := range bVals {
		if ln := val.Len(); ln > maxLen {
			maxLen = ln
		}
	}

	onesStart, _ := slices.BinarySearchFunc(bVals, BinaryRepresentation{val: 1 << (maxLen - 1)}, cmpBinary)
	ones := bVals[onesStart:]

	mx := 0

	for _, val := range ones {
		v := bestMatch(val, maxLen, bVals)
		if v > mx {
			mx = v
		}
	}

	return mx
}

func main() {

	rdr := bufio.NewReader(os.Stdin)

	readIntSlice(rdr)
	vals := readIntSlice(rdr)

	fmt.Println(maxXor(vals))

}
