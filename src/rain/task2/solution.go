package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func isComplete(directions map[int]int) bool {
	for _, v := range directions {
		if v == 0 {
			return false
		}
	}
	return true
}

func Solution(A []int) int {
	// write your code in Go 1.4
	directions := make(map[int]int)
	for _, k := range A {
		directions[k] = 0
		//if n,ok := directions[k]; ok {
		//	directions[k] = n+1
		//} else {
		//	directions[k]=0
		//}
	}
	minDays := len(A)
	j := 0
	k := 0
	for ; k < len(A); k++ {
		directions[A[k]] = directions[A[k]] + 1
		if isComplete(directions) {
			if k-j < minDays {
				minDays = k - j + 1
			}
			for ; j < k; j++ {
				directions[A[j]] = directions[A[j]] - 1
				if isComplete(directions) {
					if k-j < minDays {
						minDays = k - j
					}
				} else {
					j++
					break
				}
			}
		}
	}
	return minDays
}
