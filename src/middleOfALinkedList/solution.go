package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

func middleOfLinkedList(head *Node) int {
	if head.next == nil {
		return head.val
	}
	c := head
	for r := head; r != nil && r.next != nil; r = r.next.next {
		c = c.next
	}
	return c.val
}

func buildList(nodes []string, pos *int) *Node {
	if *pos == len(nodes) {
		return nil
	}
	val, _ := strconv.Atoi(nodes[*pos])
	*pos++
	nxt := buildList(nodes, pos)
	return &Node{val, nxt}
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
	headIdx := 0
	head := buildList(splitWords(scanner.Text()), &headIdx)
	res := middleOfLinkedList(head)
	fmt.Println(res)
}
