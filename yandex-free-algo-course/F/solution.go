package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ForRemove(a, b byte) bool {
	v := int(a) - int(b)
	if v < 0 {
		v = -v
	}
	return v == 32
}

func readString(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func refineString(arr []byte) string {

	res := make([]byte, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(res) == 0 {
			res = append(res, arr[i])
			continue
		}
		if ForRemove(res[len(res)-1], arr[i]) {
			res = res[:len(res)-1]
		} else {
			res = append(res, arr[i])
		}
	}

	return string(res)
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	str := readString(rdr)
	fmt.Println(refineString([]byte(str)))
}
