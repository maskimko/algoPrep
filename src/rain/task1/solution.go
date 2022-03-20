package main

import (
	"fmt"
	"os"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

var somePrimes []int = []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}

func nextInt(i int, limit int) int {
	return ((i+somePrimes[i%len(somePrimes)])*i - somePrimes[i%len(somePrimes)] + limit) % limit
}

func Solution(N int) []int {
	// write your code in Go 1.4
	offset := 1
	if N < 1 || N > 1000 {
		_, _ = fmt.Fprintf(os.Stderr, "wrong input")
		return nil
	}
	var result []int
	for i := 0; i < N; i++ {
		result = append(result, nextInt(N-i+offset, 100)+offset)
	}
	return result
}
