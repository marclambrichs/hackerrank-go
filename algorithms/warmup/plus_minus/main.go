package main

import "fmt"

func main() {
	var total int
	var pos, neg, zeros, nr float64
	fmt.Scanf("%d", &total)
	for i := 0; i < total; i++ {
		fmt.Scanf("%f", &nr)
		if nr > 0 {
			pos++
		} else if nr < 0 {
			neg++
		} else {
			zeros++
		}
	}
	fmt.Printf("%.6f\n", pos/float64(total))
	fmt.Printf("%.6f\n", neg/float64(total))
	fmt.Printf("%.6f\n", zeros/float64(total))

}
