package main

import (
	"fmt"
)

func scanInt(n int) []int {
	input := make([]int, n)
	buffer := make([]interface{}, n)

	for i := 0; i < n; i++ {
		buffer[i] = &input[i]
	}
	fmt.Scanln(buffer...)
	return input
}

func main() {
	var a, total int
	fmt.Scanf("%d", &a)
	ints := scanInt(a)
	for _, val := range ints {
		total += val
	}
	fmt.Printf("%d\n", total)

}
