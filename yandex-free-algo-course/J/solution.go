package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func markHits(N int, field []bool, line, col int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == line || j == col {
				field[i*N+j] = true
				continue
			}
			dx := i - line
			dy := j - col
			if dx == dy || dx == -dy {
				field[i*N+j] = true
			}
		}
	}
}

func PutFerz(N int, field []bool, path string, line int) {
	//fmt.Println(path, line)

	possible := make([]int, 0)
	for i := range N {
		if !field[line*N+i] {
			possible = append(possible, i)
		}
	}

	//fmt.Println("Possible:", possible)
	for _, p := range possible {
		pth := path + " " + strconv.Itoa(p+1)

		if line == N-1 {
			res = append(res, pth[1:])
			continue
		}

		fld := slices.Clone(field)
		markHits(N, fld, line, p)
		PutFerz(N, fld, pth, line+1)
	}
}

var res []string

func main() {
	rdr := bufio.NewReader(os.Stdin)
	N := readIntSlice(rdr)[0]

	field := make([]bool, N*N)

	PutFerz(N, field, "", 0)

	fmt.Println(len(res))
	for _, s := range res {
		fmt.Println(s)
	}
}
