package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readIntSlice(r *bufio.Reader) []int {
	line, _ := r.ReadString('\n')
	fields := strings.Fields(line)
	res := make([]int, len(fields))
	for i, f := range fields {
		rs, _ := strconv.Atoi(f)
		res[i] = rs
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func main() {

	rdr := bufio.NewReader(os.Stdin)
	readIntSlice(rdr)
	heaps_ := readIntSlice(rdr)
	power := 0
	h := IntHeap(heaps_)
	heap.Init(&h)

	for h.Len() > 1 {
		currPwr := heap.Pop(&h).(int) + heap.Pop(&h).(int)
		power += currPwr
		heap.Push(&h, currPwr)
	}

	fmt.Println(power)
}
