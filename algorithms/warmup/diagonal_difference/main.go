package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var m = make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &m[i][j])
		}
	}

	var d0, d1 int
	for i := 0; i < n; i++ {
		d0 += m[i][i]
	}
	for i := 0; i < n; i++ {
		d1 += m[i][n-i-1]
	}
	fmt.Printf("%d", int(math.Abs(float64(d0-d1))))
}
