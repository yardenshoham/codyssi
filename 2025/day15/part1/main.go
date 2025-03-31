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

func insert(root, newNode *node) {
	if newNode.artifactID > root.artifactID {
		if root.right == nil {
			root.right = newNode
			return
		}
		insert(root.right, newNode)
		return
	}
	if root.left == nil {
		root.left = newNode
		return
	}
	insert(root.left, newNode)
}

func occupied(root *node) int {
	if root == nil {
		return 0
	}
	return 1 + max(occupied(root.left), occupied(root.right))
}

func layerSum(root *node, n int) int {
	if root == nil {
		return 0
	}
	if n == 0 {
		return root.artifactID
	}
	return layerSum(root.left, n-1) + layerSum(root.right, n-1)
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
		insert(root, &n)
	}
	layers := occupied(root)
	maxSum := 0
	for i := range layers {
		maxSum = max(maxSum, layerSum(root, i))
	}
	fmt.Println(layers * maxSum)
}
