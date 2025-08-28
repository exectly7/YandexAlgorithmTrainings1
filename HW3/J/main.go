package main

import (
	"fmt"
)

type rect struct {
	umin, umax int
	vmin, vmax int
}

func main() {
	var t, d, n int
	fmt.Scan(&t, &d, &n)

	prev := rect{0, 0, 0, 0}

	var x, y int
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)

		prev.umin -= t
		prev.umax += t
		prev.vmin -= t
		prev.vmax += t

		u := x + y
		v := x - y
		ps := rect{u - d, u + d, v - d, v + d}

		prev = intersect(prev, ps)
		if prev.umin > prev.umax || prev.vmin > prev.vmax {
			fmt.Println(0)
			return
		}
	}

	points := make([][2]int, 0)
	for u := prev.umin; u <= prev.umax; u++ {
		for v := prev.vmin; v <= prev.vmax; v++ {
			if (u+v)%2 != 0 {
				continue
			}
			x := (u + v) / 2
			y := (u - v) / 2
			if abs(x+y-u) <= 0 && abs(x-y-v) <= 0 {
				points = append(points, [2]int{x, y})
			}
		}
	}

	fmt.Println(len(points))
	for _, p := range points {
		fmt.Println(p[0], p[1])
	}
}

func intersect(a, b rect) rect {
	return rect{
		umin: max(a.umin, b.umin),
		umax: min(a.umax, b.umax),
		vmin: max(a.vmin, b.vmin),
		vmax: min(a.vmax, b.vmax),
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
