package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	artifactName string
	artifactID   int
	left, right  *node
}

func insert(root, newNode *node, seq *[]string) {
	*seq = append(*seq, root.artifactName)
	if newNode.artifactID > root.artifactID {
		if root.right == nil {
			root.right = newNode
			return
		}
		insert(root.right, newNode, seq)
		return
	}
	if root.left == nil {
		root.left = newNode
		return
	}
	insert(root.left, newNode, seq)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(input), "\n\n")
	if len(parts) != 2 {
		panic("should have 2 parts")
	}
	var root *node
	for _, line := range strings.Split(parts[0], "\n") {
		var n node
		_, err := fmt.Sscanf(line, "%s | %d", &n.artifactName, &n.artifactID)
		if err != nil {
			panic(err)
		}
		if root == nil {
			root = &n
			continue
		}
		insert(root, &n, &[]string{})
	}
	seq := []string{}
	insert(root, &node{artifactID: 500000}, &seq)
	fmt.Println(strings.Join(seq, "-"))
}
