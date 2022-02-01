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
	result := make([][]int, 1)
	result[0] = append(result[0], root.val)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	bfs(&queue, &result)
	return result
}

func bfs(queue *[]*Node, result *[][]int) {
	level := 0
	for len(*queue) > 0 {

		items := len(*queue)
		var levelQueue []*Node
		level++
		for i := 0; i < items; i++ {
			node := (*queue)[0]
			*queue = (*queue)[1:]
			if node != nil {
				if node.left != nil {
					levelQueue = append(levelQueue, node.left)
				}
				if node.right != nil {
					levelQueue = append(levelQueue, node.right)
				}
			}
		}
		if len(levelQueue) > 0 {
			*result = append(*result, nil)
			for _, n := range levelQueue {
				*queue = append(*queue, n)
			}
			if level%2 == 0 {
				for _, n := range levelQueue {
					(*result)[level] = append((*result)[level], n.val)
				}
			} else {
				for n := len(levelQueue) - 1; n >= 0; n-- {
					(*result)[level] = append((*result)[level], levelQueue[n].val)
				}
			}
		}
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
