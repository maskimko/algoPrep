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

func levelOrderTraversal(root *Node) [][]int {
	result := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	bfs(&queue, &result)
	return result
}

func bfs(queue *[]*Node, result *[][]int) {
	level := 0
	for len(*queue) > 0 {
		n := len(*queue)
		*result = append(*result, nil)
		for i := 0; i < n; i++ {
			node := (*queue)[0]
			*queue = (*queue)[1:]
			(*result)[level] = append((*result)[level], node.val)

			if node.left != nil {
				*queue = append(*queue, node.left)
			}
			if node.right != nil {
				*queue = append(*queue, node.right)
			}
		}
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
	res := levelOrderTraversal(root)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
