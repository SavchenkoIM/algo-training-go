package main

import (
	"bufio"
	"fmt"
	"os"
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

type Node struct {
	val     int
	maxPath int
	next    []*Node
}

func MaxPathFromNode(root *Node, ln int) int {

	if root.maxPath > 0 {
		return root.maxPath + ln - 1
	}
	if root.next == nil {
		root.maxPath = 1
		return ln
	}
	maxLen := ln
	for _, n := range root.next {
		if mpn := MaxPathFromNode(n, ln+1); mpn > maxLen {
			maxLen = mpn
		}
	}

	if root.maxPath < maxLen-ln+1 {
		root.maxPath = maxLen - ln + 1
	}
	return maxLen
}

func main() {

	rdr := bufio.NewReader(os.Stdin)

	dim := readIntSlice(rdr)
	matrix := make([][]int, 0, dim[0])
	for range dim[0] {
		matrix = append(matrix, readIntSlice(rdr))
	}

	nodeMatrix := make([][]*Node, dim[0])
	for i := range len(nodeMatrix) {
		nodeMatrix[i] = make([]*Node, dim[1])
	}
	for y := 0; y < dim[0]; y++ {
		for x := 0; x < dim[1]; x++ {
			nodeMatrix[y][x] = &Node{val: matrix[y][x]}
		}
	}

	for y := 0; y < dim[0]; y++ {
		for x := 0; x < dim[1]; x++ {
			if x > 0 && matrix[y][x] < matrix[y][x-1] {
				nodeMatrix[y][x].next = append(nodeMatrix[y][x].next, nodeMatrix[y][x-1])
			}
			if y > 0 && matrix[y][x] < matrix[y-1][x] {
				nodeMatrix[y][x].next = append(nodeMatrix[y][x].next, nodeMatrix[y-1][x])
			}
			if x < dim[1]-1 && matrix[y][x] < matrix[y][x+1] {
				nodeMatrix[y][x].next = append(nodeMatrix[y][x].next, nodeMatrix[y][x+1])
			}
			if y < dim[0]-1 && matrix[y][x] < matrix[y+1][x] {
				nodeMatrix[y][x].next = append(nodeMatrix[y][x].next, nodeMatrix[y+1][x])
			}
		}
	}

	maxLen := 0
	for y := 0; y < dim[0]; y++ {
		for x := 0; x < dim[1]; x++ {
			if mpn := MaxPathFromNode(nodeMatrix[y][x], 1); mpn > maxLen {
				maxLen = mpn
			}
		}
	}

	fmt.Println(maxLen)
}
