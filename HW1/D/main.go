package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	if a == 0 {
		if c*c == b {
			fmt.Println("MANY SOLUTIONS")
		} else {
			fmt.Println("NO SOLUTION")
		}
	} else {
		if (c*c-b)%a == 0 {
			fmt.Println((c*c - b) / a)
		} else {
			fmt.Println("NO SOLUTION")
		}
	}
}
