package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readString(rdr *bufio.Reader) string {
	j, _ := rdr.ReadString('\n')
	return strings.Trim(j, "\r\n")
}

func readIntSlice(rdr *bufio.Reader) []int {
	s := readString(rdr)
	res := make([]int, 0)
	for _, i := range strings.Fields(s) {
		ii, _ := strconv.Atoi(i)
		res = append(res, ii)
	}
	return res
}

type Town struct {
	Id, X, Y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Dest(p1, p2 Town) int {
	return abs(p1.Y-p2.Y) + abs(p1.X-p2.X)
}

func isHopPossible(p1, p2 Town, rng int) bool {
	return Dest(p1, p2) <= rng && p1.Id != p2.Id
}

// 0 - self
// >0 - Possible
// -1 - Impossible
// -2 - Not calculated

func cmpByX(t1, t2 Town) int {
	if t1.X > t2.X {
		return 1
	}
	if t1.X < t2.X {
		return -1
	}
	return 0
}

func cmpByY(t1, t2 Town) int {
	if t1.Y > t2.Y {
		return 1
	}
	if t1.Y < t2.Y {
		return -1
	}
	return 0
}

func getHopableTowns(t Town, rng int, tbx, tby []Town) []int {
	res := make([]int, 0)

	byX := make(map[int]struct{})
	xS, _ := slices.BinarySearchFunc(tbx, Town{X: t.X - rng}, cmpByX)
	xF, _ := slices.BinarySearchFunc(tbx, Town{X: t.X + rng + 1}, cmpByX)
	for i := xS; i < xF; i++ {
		if isHopPossible(t, tbx[i], rng) {
			byX[tbx[i].Id] = struct{}{}
		}
	}

	yS, _ := slices.BinarySearchFunc(tby, Town{Y: t.Y - rng}, cmpByY)
	yF, _ := slices.BinarySearchFunc(tby, Town{Y: t.Y + rng + 1}, cmpByY)
	for i := yS; i < yF; i++ {
		if isHopPossible(t, tby[i], rng) {
			if _, ok := byX[tby[i].Id]; ok {
				res = append(res, tby[i].Id)
			}
		}
	}

	slices.Sort(res)
	return res
}

func main() {

	rdr := bufio.NewReader(os.Stdin)
	townsN := readIntSlice(rdr)[0]
	towns := make([]Town, 0)

	for t := 0; t < townsN; t++ {
		crd := readIntSlice(rdr)
		towns = append(towns, Town{Id: t, X: crd[0], Y: crd[1]})
	}
	carRange := readIntSlice(rdr)[0]
	destCrd := readIntSlice(rdr)
	from, to := destCrd[0]-1, destCrd[1]-1

	townsByX := slices.Clone(towns)
	slices.SortFunc(townsByX, cmpByX)
	townsByY := slices.Clone(towns)
	slices.SortFunc(townsByY, cmpByY)

	moved := true
	Visited := map[int]bool{from: true} // TownId: Visited
	Routes := map[int]int{from: 0}      // TownId: PathLen
	for moved {
		moved = false
		currRoutes := map[int]int{}
		for k, ln := range Routes {
			fwd := getHopableTowns(towns[k], carRange, townsByX, townsByY)
			for _, next := range fwd {
				if !Visited[next] {
					Visited[next] = true
					currRoutes[next] = ln + 1
					moved = true

					if next == to {
						fmt.Println(ln + 1)
						return
					}
				}
			}
		}
		Routes = currRoutes
	}

	fmt.Println("-1")
}
