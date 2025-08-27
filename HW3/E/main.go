package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	numbers := make(map[int]struct{})
	numbers[a] = struct{}{}
	numbers[b] = struct{}{}
	numbers[c] = struct{}{}
	var num int
	fmt.Scan(&num)
	rem := 0
	ct := 0
	if num == 0 {
		_, ok := numbers[0]
		if ok {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
		return
	}
	for ; num > 0; num /= 10 {
		rem = num % 10
		if _, ok := numbers[rem]; !ok {
			numbers[rem] = struct{}{}
			ct++
		}
	}

	fmt.Println(ct)
}
