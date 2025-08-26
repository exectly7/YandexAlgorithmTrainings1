package main

import "fmt"

func main() {
	var a1, a2, b1, b2 int
	fmt.Scan(&a1, &a2, &b1, &b2)

	s1 := (a1 + b1) * max(a2, b2)
	s2 := (a1 + b2) * max(a2, b1)
	s3 := (a2 + b1) * max(a1, b2)
	s4 := (a2 + b2) * max(a1, b1)

	if s1 <= s2 && s1 <= s3 && s1 <= s4 {
		fmt.Println(a1+b1, max(a2, b2))
	} else if s2 <= s1 && s2 <= s3 && s2 <= s4 {
		fmt.Println(a1+b2, max(a2, b1))
	} else if s3 <= s1 && s3 <= s2 && s3 <= s4 {
		fmt.Println(a2+b1, max(a1, b2))
	} else {
		fmt.Println(a2+b2, max(a1, b1))
	}
}
