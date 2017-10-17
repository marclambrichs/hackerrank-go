package main

import "fmt"

func scanInt(n int) []int {
	input := make([]int, n)
	buffer := make([]interface{}, n)

	for i := 0; i < n; i++ {
		buffer[i] = &input[i]
	}
	fmt.Scanln(buffer...)
	return inputva
}

func main() {
	var alice, bob int
	a := scanInt(3)
	b := scanInt(3)
	for i, _ := range a {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	fmt.Printf("%d %d\n", alice, bob)
}
