package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pq []int

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	i := x.(int)
	*p = append(*p, i)
}

func (p *pq) Pop() interface{} {
	c := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return c
}

func heapTop3(arr []int) []int {
	h := make(pq, 0)
	heap.Init(&h)
	for i := 0; i < len(arr); i++ {
		heap.Push(&h, arr[i])
		if len(h) > 3 {
			heap.Pop(&h)
		}
	}
	sort.Ints(h)
	return h
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
	arr := arrayAtoi(splitWords(scanner.Text()))
	res := heapTop3(arr)
	fmt.Println(strings.Join(arrayItoa(res), " "))
}
