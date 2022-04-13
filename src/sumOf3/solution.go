package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func tripletsWithSum0(nums []int) [][]int {
	data := make([]int, len(nums), len(nums))
	copy(data, nums)
	sort.Ints(data)
	var result [][]int

	memo := make(map[int][]int)
	for i := 0; i < len(data)-2; i++ {
		k := i + 1
		for j := len(data) - 1; j > k; {
			if data[k]+data[j] < -data[i] {
				k++
			} else {
				if data[i]+data[k]+data[j] == 0 {

					key := data[i] ^ data[k] ^ data[j]
					candidate := []int{data[i], data[k], data[j]}
					if previous, ok := memo[key]; ok {
						same := true
						for n := 0; n < len(previous); n++ {
							if previous[n] != candidate[n] {
								same = false
								break
							}
						}
						if same {
							k++
							j++
							continue
						}
					}
					result = append(result, []int{data[i], data[k], data[j]})
					memo[key] = []int{data[i], data[k], data[j]}
					k++
					j--
				} else {
					j--
				}
			}
		}

	}
	return result
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func arrayItoa(arr []int) []string {
	res := []string{}
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := arrayAtoi(splitWords(scanner.Text()))
	res := tripletsWithSum0(nums)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
