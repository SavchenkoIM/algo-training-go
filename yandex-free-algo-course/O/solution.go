package main

import (
	"bufio"
	"container/heap"
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

type Building struct {
	Price int
	Cap   int
}

func cmpByPrice(a, b Building) int {
	if a.Price > b.Price {
		return 1
	}
	if a.Price < b.Price {
		return -1
	}
	return 0
}

func cmpByCapInv(a, b Building) int {
	if a.Cap > b.Cap {
		return -1
	}
	if a.Cap < b.Cap {
		return 1
	}
	return 0
}

// Priority queue implementation
// An IntHeap is a MAX-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Priority queue implementation

func main() {

	rdr := bufio.NewReader(os.Stdin)

	mData := readIntSlice(rdr)
	N, canBuy := mData[0], mData[1]

	blds := make([]Building, N)
	for i := range N {
		cb := readIntSlice(rdr)
		blds[i] = Building{Price: cb[0], Cap: cb[1]}
	}

	ccap := readIntSlice(rdr)[0]
	slices.SortFunc(blds, cmpByCapInv)
	slices.SortStableFunc(blds, cmpByPrice)

	var awailCaps IntHeap
	heap.Init(&awailCaps)

	for range canBuy {
		endIndex, _ := slices.BinarySearchFunc(blds, Building{Price: ccap + 1}, cmpByPrice)
		endIndex--

		for i := 0; i <= endIndex; i++ {
			heap.Push(&awailCaps, blds[i].Cap)
		}

		if len(awailCaps) == 0 {
			break
		}

		ccap += heap.Pop(&awailCaps).(int)

		blds = blds[endIndex+1:]
	}

	fmt.Println(ccap)
}
