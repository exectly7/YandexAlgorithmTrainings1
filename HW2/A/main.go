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

func solve(a []int) string {
	if len(a) == 0 {
		return "YES"
	}
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] {
			return "NO"
		}
	}
	return "YES"
}
