package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func binaryTreeRightSideView(root *Node) []int {
	return rSide([]*Node{root})
}

func rSide(q []*Node) []int {
	var results []int
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			if i == levelSize-1 {
				results = append(results, q[i].val)
			}
			if q[i].left != nil {
				q = append(q, q[i].left)
			}
			if q[i].right != nil {
				q = append(q, q[i].right)
			}
		}
		q = q[levelSize:]
	}
	return results
}

func buildTree(nodes []string, pos *int) *Node {
	val := nodes[*pos]
	*pos++
	if val == "x" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	left := buildTree(nodes, pos)
	right := buildTree(nodes, pos)
	return &Node{v, left, right}
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
	rootIdx := 0
	root := buildTree(splitWords(scanner.Text()), &rootIdx)
	res := binaryTreeRightSideView(root)
	fmt.Println(strings.Join(arrayItoa(res), " "))
}
