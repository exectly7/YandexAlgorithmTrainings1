package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	minABC := min(a, b, c)
	midABC := a + b + c - minABC - max(a, b, c)
	if max(d, e) >= max(midABC, minABC) && min(d, e) >= min(midABC, minABC) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
