package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func squareRoot(n int) int {
    var left, right int
    if n == 0 {
        return 0
    }
    for right = n; left<=right; {
        mid := left + (right-left)/2
        square := mid*mid
        if square == n {
            return mid
        } else if (mid*mid) >n {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return left
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    res := squareRoot(n)
    fmt.Println(res)
}