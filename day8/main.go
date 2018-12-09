package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	metadata []int
}

func parseNode(parts []int) (*node, int) {
	root := &node{}
	nChildren := parts[0]
	nMetadata := parts[1]

	dataParts := parts[2:]
	for i := 0; i < nChildren; i++ {
		child, partsLen := parseNode(dataParts)
		dataParts = dataParts[partsLen:]

		root.children = append(root.children, child)
	}

	for i := 0; i < nMetadata; i++ {
		root.metadata = append(root.metadata, dataParts[i])
	}
	dataParts = dataParts[nMetadata:]

	return root, len(parts) - len(dataParts)
}

func (n *node) traverse(f func(n *node)) {
	f(n)

	for _, v := range n.children {
		v.traverse(f)
	}
}

func (n *node) value() int {
	var sum int

	if len(n.children) == 0 {
		n.traverse(func(n *node) {
			for _, v := range n.metadata {
				sum += v
			}
		})
		return sum
	}

	for _, ref := range n.metadata {
		idx := ref - 1
		if idx >= 0 && idx < len(n.children) {
			sum += n.children[idx].value()
		}
	}

	return sum
}

func main() {
	var parts []int

	f, _ := ioutil.ReadFile("input.txt")

	for _, p := range strings.Split(strings.Trim(string(f), "\n"), " ") {
		part, _ := strconv.Atoi(p)
		parts = append(parts, part)
	}

	root, _ := parseNode(parts)

	var sum int
	root.traverse(func(n *node) {
		for _, v := range n.metadata {
			sum += v
		}
	})

	fmt.Println("sum of all metadata entries:", sum)

	fmt.Println("value of the root node:", root.value())
}
