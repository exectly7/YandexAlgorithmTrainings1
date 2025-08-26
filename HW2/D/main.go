package main

import (
	"fmt"
)

func main() {
	var num int
	var numbers []int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}

	fmt.Println(solve(numbers))
}

func solve(a []int) int {
	if len(a) < 3 {
		return 0
	}
	c := 0
	for i := 1; i < len(a)-1; i++ {
		if a[i] > a[i+1] && a[i] > a[i-1] {
			c++
		}
	}
	return c
}
