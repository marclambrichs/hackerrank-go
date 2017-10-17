package main

import "fmt"

func main() {
	var count, n, height int
	max := 0
	fmt.Scanf("%v\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%v", &height)
		if height == max {
			count++
		} else if height > max {
			max = height
			count = 1
		}
	}
	fmt.Println(count)
}
