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

func binaryTreeMinDepth(root *Node) int {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	return bfsMinDepth(queue)
}

func bfsMinDepth(q []*Node) int {
	depth := 0
	for len(q) > 0 {

		n := len(q)
		for i := 0; i < n; i++ {
			if q[i] == nil {
				return depth - 1 //Should never happer
			}
			if q[i].left == nil && q[i].right == nil {
				return depth
			}
			if q[i].left != nil {
				q = append(q, q[i].left)
			}
			if q[i].right != nil {
				q = append(q, q[i].right)
			}
		}
		depth++
		q = q[n:]
	}
	return depth
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rootIdx := 0
	root := buildTree(splitWords(scanner.Text()), &rootIdx)
	res := binaryTreeMinDepth(root)
	fmt.Println(res)
}
