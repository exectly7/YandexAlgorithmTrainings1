package main

import (
	"fmt"
	"math"
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

func solve(a []int) (int, int, int) {
	if len(a) == 3 {
		return a[0], a[1], a[2]
	}
	mxPos1 := math.MinInt
	mxPos2 := math.MinInt
	mxPos3 := math.MinInt
	mnNeg1 := math.MaxInt
	mnNeg2 := math.MaxInt
	allNeg := true
	for _, n := range a {
		if n > 0 {
			allNeg = false
			if n > mxPos1 {
				mxPos3 = mxPos2
				mxPos2 = mxPos1
				mxPos1 = n
			} else if n > mxPos2 {
				mxPos3 = mxPos2
				mxPos2 = n
			} else if n > mxPos3 {
				mxPos3 = n
			}
		} else {
			if n < mnNeg1 {
				mnNeg2 = mnNeg1
				mnNeg1 = n
			} else if n < mnNeg2 {
				mnNeg2 = n
			}
		}
	}
	if allNeg {
		mxPos1 = math.MinInt
		mxPos2 = math.MinInt
		mxPos3 := math.MaxInt
		for _, n := range a {
			if n > mxPos1 {
				mxPos3 = mxPos2
				mxPos2 = mxPos1
				mxPos1 = n
			} else if n > mxPos2 {
				mxPos3 = mxPos2
				mxPos2 = n
			} else if n > mxPos3 {
				mxPos3 = n
			}
		}
		return mxPos3, mxPos2, mxPos1
	} else {
		if mnNeg1*mnNeg2*mxPos1 < mxPos1*mxPos2*mxPos3 && mxPos3 != math.MinInt {
			return mxPos3, mxPos2, mxPos1
		} else {
			return mnNeg1, mnNeg2, mxPos1
		}
	}
}
