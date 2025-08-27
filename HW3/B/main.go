package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	line1, _ := in.ReadString('\n')
	line2, _ := in.ReadString('\n')

	nums1 := parseLine(line1)
	nums2 := parseLine(line2)

	result := solve(nums1, nums2)

	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func parseLine(line string) map[int]struct{} {
	parts := strings.Fields(line)
	nums := make(map[int]struct{})
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		nums[n] = struct{}{}
	}
	return nums
}

func solve(a map[int]struct{}, b map[int]struct{}) (res []int) {
	for n := range a {
		if _, ok := b[n]; ok {
			res = append(res, n)
		}
	}
	sort.Ints(res)
	return
}
