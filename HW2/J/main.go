package main

import "fmt"

func main() {
	var n int
	var prev float64
	var now float64
	var str string
	fmt.Scan(&n)
	fmt.Scan(&prev)
	mx := 4000.0
	mn := 30.0
	for i := 0; i < n-1; i++ {
		fmt.Scan(&now, &str)
		if str == "closer" {
			if now > prev {
				if (prev+now)/2 > mn {
					mn = (prev + now) / 2
				}
			} else {
				if (prev+now)/2 < mx {
					mx = (prev + now) / 2
				}
			}
		} else {
			if now > prev {
				if (prev+now)/2 < mx {
					mx = (prev + now) / 2
				}
			} else {
				if (prev+now)/2 > mn {
					mn = (prev + now) / 2
				}
			}
		}
		prev = now
	}
	fmt.Println(mn, mx)
}

func solve() {

}
