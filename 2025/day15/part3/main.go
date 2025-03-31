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

func find(root, newNode *node, seq *[]string) {
	*seq = append(*seq, root.artifactName)
	if newNode.artifactID > root.artifactID {
		if root.right == nil {
			return
		}
		find(root.right, newNode, seq)
		return
	}
	if root.left == nil {
		return
	}
	find(root.left, newNode, seq)
}

func parseLine(line string) node {
	var n node
	_, err := fmt.Sscanf(line, "%s | %d", &n.artifactName, &n.artifactID)
	if err != nil {
		panic(err)
	}
	return n
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
		n := parseLine(line)
		if root == nil {
			root = &n
			continue
		}
		insert(root, &n, &[]string{})
	}
	lastLines := strings.Split(parts[1], "\n")
	if len(lastLines) != 2 {
		panic("should have 2 lines")
	}
	n0 := parseLine(lastLines[0])
	n1 := parseLine(lastLines[1])
	seq0 := []string{}
	seq1 := []string{}
	find(root, &n0, &seq0)
	find(root, &n1, &seq1)
	var ancestor string
	for i := range min(len(seq0), len(seq1)) {
		if seq0[i] != seq1[i] {
			break
		}
		ancestor = seq0[i]
	}
	fmt.Println(ancestor)
}
