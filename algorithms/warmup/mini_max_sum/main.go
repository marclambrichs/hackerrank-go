package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func scanInt(s string) []int64 {
	var n []int64
	for _, f := range strings.Fields(s) {
		i, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			panic(err)
		}
		n = append(n, i)
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		arr := scanInt(scanner.Text())
		min := big.NewInt(int64(0))
		max := big.NewInt(int64(0))
		total := new(big.Int)
		for idx, _ := range arr {
			total = big.NewInt(int64(0))
			for idx2, value := range arr {
				if idx != idx2 {
					total = total.Add(total, big.NewInt(value))

				}

			}
			if total.Cmp(max) > 0 {
				max = total
			}
			if min.Cmp(big.NewInt(0)) == 0 || total.Cmp(min) < 0 {
				min = total
			}
		}
		fmt.Printf("%s %s\n", min, max)
	}
}
