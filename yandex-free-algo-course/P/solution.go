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

func readIntSlice(buf *bufio.Reader) []int32 {
	str := readString(buf)
	fields := strings.Fields(str)
	res := make([]int32, 0, len(fields))
	for _, f := range fields {
		i, _ := strconv.Atoi(f)
		res = append(res, int32(i))
	}
	return res
}

type Node struct {
	id     int32
	left   int32
	right  int32
	isLeaf bool
}

var nodes []Node

func printLevels(currLevel []Node) {

	for len(currLevel) > 0 {

		for i, nd := range currLevel {
			if i == 0 || i == len(currLevel)-1 || nd.isLeaf {
				fmt.Printf("%d ", nd.id)
			}
		}

		nextLevel := make([]Node, 0)
		for _, nd := range currLevel {
			if nd.left != -1 {
				nextLevel = append(nextLevel, nodes[nd.left])
			}
			if nd.right != -1 {
				nextLevel = append(nextLevel, nodes[nd.right])
			}
		}
		currLevel = nextLevel

	}

}

func main() {

	rdr := bufio.NewReader(os.Stdin)

	mData := readIntSlice(rdr)
	N, rootId := mData[0], mData[1]

	for i := 0; i < int(N); i++ {
		children := readIntSlice(rdr)
		nodes = append(nodes, Node{id: int32(i), left: children[0], right: children[1], isLeaf: children[0] == -1 && children[1] == -1})
	}

	printLevels([]Node{nodes[rootId]})
}
