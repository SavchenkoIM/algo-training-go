package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type ValueTag struct {
	i         int
	factor200 int
}

func CmpTags(a ValueTag, b ValueTag) int {
	if a.factor200 > b.factor200 {
		return 1
	}
	if a.factor200 < b.factor200 {
		return -1
	}
	return 0
}

func CmpTagsInv(a ValueTag, b ValueTag) int {
	if a.factor200 < b.factor200 {
		return 1
	}
	if a.factor200 > b.factor200 {
		return -1
	}
	return 0
}

var startIndex = 0
var stopIndex = 0
var prevFactor = -9999999999999999

func getNumPairsForI(currI ValueTag, elems []ValueTag, elemsInv []ValueTag) int {
	var startFound, stopFound bool

	if currI.factor200 != prevFactor {
		startIndex, startFound = slices.BinarySearchFunc(elems, currI, CmpTags)
		stopIndex, stopFound = slices.BinarySearchFunc(elemsInv, currI, CmpTagsInv)
		if stopFound {
			stopIndex = len(elemsInv) - stopIndex
		} else {
			stopIndex = len(elemsInv)
		}
	} else {
		startIndex++
		startFound = true
	}
	prevFactor = currI.factor200

	if !startFound {
		return 0
	}

	dec := 0
	idx := slices.IndexFunc(elems[startIndex:stopIndex], func(tag ValueTag) bool {
		return tag.i == currI.i
	})
	if idx >= 0 {
		dec = -1
	}

	return stopIndex - startIndex + dec
}

func getNumPairs(arr []int) int {

	if len(arr) < 2 {
		return 0
	}

	elems := make([]ValueTag, len(arr))
	for i := 0; i < len(arr); i++ {
		elems[i] = ValueTag{i, arr[i] % 200}
	}
	slices.SortStableFunc(elems, CmpTags)

	elemsInv := slices.Clone(elems)
	slices.Reverse(elemsInv)

	res := 0
	for i := 0; i < len(elemsInv); i++ {
		res += getNumPairsForI(elems[i], elems, elemsInv)
	}

	return res
}

func main() {

	in := bufio.NewReader(os.Stdin)

	scanInts(in)
	arr := scanInts(in)

	fmt.Println(getNumPairs(arr))

}
