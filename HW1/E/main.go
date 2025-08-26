package main

import (
	"fmt"
)

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)

	if n2 > m || n2 < 1 || p2 < 1 {
		fmt.Println(-1, -1)
		return
	}

	p1, n1 := -2, -2
	found := false

	for x := 1; x <= k2; x++ {
		entr := m * x

		low := (p2-1)*entr + (n2-1)*x + 1
		high := (p2-1)*entr + n2*x
		if k2 < low || k2 > high {
			continue
		}

		p := (k1-1)/entr + 1
		n := ((k1-1)%entr)/x + 1

		if !found {
			p1, n1 = p, n
			found = true
		} else {
			if p1 != p {
				p1 = 0
			}
			if n1 != n {
				n1 = 0
			}
		}
	}

	if !found {
		fmt.Println(-1, -1)
	} else {
		fmt.Println(p1, n1)
	}
}
