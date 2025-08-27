package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var input int
	anya := make(map[int]struct{})
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		anya[input] = struct{}{}
	}
	borya := make(map[int]struct{})
	for i := 0; i < m; i++ {
		fmt.Scan(&input)
		borya[input] = struct{}{}
	}
	cross := make(map[int]struct{})
	for v, _ := range anya {
		if _, ok := borya[v]; ok {
			cross[v] = struct{}{}
			delete(borya, v)
			delete(anya, v)
		}
	}
	a := mapToSortedArray(anya)
	b := mapToSortedArray(borya)
	c := mapToSortedArray(cross)
	fmt.Println(len(c))
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println(len(b))
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}

func mapToSortedArray(m map[int]struct{}) (res []int) {
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return
}
