package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	var amK, counter int
	left := k % m
	if m > k {
		fmt.Println(0)
		return
	}
	for n >= k {
		amK = n / k
		n = n % k
		counter += amK * (k / m)
		n += left * amK
	}
	fmt.Println(counter)
}
