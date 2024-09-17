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

type Point struct {
	x, y int
}

func (p Point) isOnLine(start, end Point) bool {
	dy1 := p.y - start.y
	dx1 := p.x - start.x
	dy2 := end.y - p.y
	dx2 := end.x - p.x
	if dy1 == 0 && dy2 == 0 || dy1 == 0 && p == start || dy2 == 0 && p == end {
		return true
	}
	return dy1*dx2 == dy2*dx1
}

func main() {

	rdr := bufio.NewReader(os.Stdin)

	N := readIntSlice(rdr)[0]
	pts := make([]Point, N)

	maxX := -9999999999999
	idMaxX := -1
	minX := 9999999999999
	idMinX := -1

	for i := 0; i < N; i++ {
		pt := readIntSlice(rdr)
		pts[i] = Point{x: pt[0], y: pt[1]}

		if pts[i].x > maxX {
			maxX = pts[i].x
			idMaxX = i
		}

		if pts[i].x < minX {
			minX = pts[i].x
			idMinX = i
		}
	}

	for _, pt := range pts {
		if !pt.isOnLine(pts[idMinX], pts[idMaxX]) {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
