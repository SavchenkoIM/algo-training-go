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

var primes []int

func isPrime(num int) bool {
	for i, prime := range primes {
		if i == 0 {
			continue
		}
		if num%prime == 0 {
			return false
		}
	}
	return true
}

func genPrimes() {
	primes = append(primes, 1, 2)
	for i := 3; i < 10000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
}

func getAllDividers(number int) []int {
	if number == 1 {
		return []int{}
	}

	var parts []int

	for i := 1; i < len(primes) && primes[i] <= number; i++ {
		for number%primes[i] == 0 {
			number = number / primes[i]
			parts = append(parts, primes[i])
		}
	}

	resMap := make(map[int]struct{})

	for i := 1; i < 1<<len(parts)-1; i++ {
		num := i
		currRes := 1
		for ptr := 0; num > 0; ptr++ {
			if num&1 == 1 {
				currRes *= parts[ptr]
			}
			num >>= 1
		}
		resMap[currRes] = struct{}{}
	}

	res := make([]int, 0, len(resMap))
	res = append(res, 1)
	for k, _ := range resMap {
		res = append(res, k)
	}

	slices.Sort(res) // Not necessary

	return res
}

type NodeType int

const (
	TYPE_UNKNOWN = NodeType(iota)
	TYPE_WIN
	TYPE_LOSE
)

type Node struct {
	number int
	next   []*Node
	Type   NodeType
}

var Index map[int]*Node = make(map[int]*Node)

func main() {
	// Generating primes
	genPrimes()

	// Creating indexed tree of possible Turns for every number <= 10000
	head := &Node{number: 10000}
	Index = map[int]*Node{10000: head}

	for i := 10000; i > 0; i-- {
		divs := getAllDividers(i)
		for _, div := range divs {
			node, ok := Index[i-div]
			if !ok {
				Index[i-div] = &Node{number: i - div}
				node = Index[i-div]
			}
			Index[i].next = append(Index[i].next, node)
		}
	}

	// Setting node types
	for i := 1; i <= 10000; i++ {
		if Index[i].next == nil {
			Index[i].Type = TYPE_LOSE
			continue
		}
		for _, nextNode := range Index[i].next {
			Index[i].Type = TYPE_LOSE
			if nextNode.Type == TYPE_LOSE {
				Index[i].Type = TYPE_WIN
				break
			}
		}
	}

	// Solve problem
	rdr := bufio.NewReader(os.Stdin)

	if Index[readIntSlice(rdr)[0]].Type == TYPE_WIN {
		fmt.Println("Pasha")
	} else {
		fmt.Println("Mark")
	}
}
