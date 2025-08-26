package main

import "fmt"

func main() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	minBy1 := n*(a+1) - a
	minBy2 := m*(b+1) - b
	maxBy1 := n*(a+1) + a
	maxBy2 := m*(b+1) + b

	if minBy1 > maxBy2 || minBy2 > maxBy1 {
		fmt.Println(-1)
	} else {
		fmt.Println(max(minBy1, minBy2), min(maxBy1, maxBy2))
	}
}
