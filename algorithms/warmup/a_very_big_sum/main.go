package main

import (
	"fmt"
	"math/big"
)

func scanInt(n int) []int64 {
	input := make([]int64, n)
	buffer := make([]interface{}, n)

	for i := 0; i < n; i++ {
		buffer[i] = &input[i]
	}
	fmt.Scanln(buffer...)
	return input
}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	nrs := scanInt(a)

	total := big.NewInt(int64(0))
	for _, x := range nrs {
		total = total.Add(total, big.NewInt(x))
	}
	fmt.Print(total)
}
