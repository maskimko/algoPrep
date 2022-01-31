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

type NodeLevel struct {
	node  *Node
	level int
}

func levelOrderTraversal(root *Node) [][]int {
	result := make([][]int, 1)
	queue := make([]NodeLevel, 0)
	queue = append(queue, NodeLevel{root, 0})
	bfs(&queue, &result)
	return result
}

func bfs(queue *[]NodeLevel, result *[][]int) {
	for level := 0; len(*queue) > 0; {
		nl := (*queue)[0]
		if nl.level > level {
			*result = append(*result, nil)
		}
		level = nl.level
		*queue = (*queue)[1:]
		(*result)[nl.level] = append((*result)[nl.level], nl.node.val)

		if nl.node.left != nil {
			*queue = append(*queue, NodeLevel{nl.node.left, nl.level + 1})
		}
		if nl.node.right != nil {
			*queue = append(*queue, NodeLevel{nl.node.right, nl.level + 1})
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
	res := levelOrderTraversal(root)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
