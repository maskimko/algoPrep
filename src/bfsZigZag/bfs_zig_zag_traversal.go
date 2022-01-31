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

func zigZagTraversal(root *Node) [][]int {
	result := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	bfs(&queue, &result)
	return result
}

func bfs(queue *[]*Node, result *[][]int) {
	level := 0
	var stack []*Node
	for len(*queue) > 0 || len(stack) > 0 {
		n := len(*queue)
		*result = append(*result, nil)
		for i := 0; i < n; i++ {
			node := (*queue)[0]
			*queue = (*queue)[1:]
			(*result)[level] = append((*result)[level], node.val)
			if level%2 == 0 {
				if node.right != nil {
					*queue = append(*queue, node.right)
				}
				if node.left != nil {
					*queue = append(*queue, node.left)
				}
			} else {
				if node.right != nil {
					stack = append([]*Node{node.right}, stack...)
				}
				if node.left != nil {
					stack = append([]*Node{node.left}, stack...)
				}

			}
		}
		for _, node := range stack {
			if node != nil {
				*queue = append(*queue, node)
			}
		}
		stack = nil
		level++
	}
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
	res := zigZagTraversal(root)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
