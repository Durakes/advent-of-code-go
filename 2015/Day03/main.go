package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	x int
	y int
}

func main() {
	input = strings.TrimRight(input, "\n")
	part1(input)
	part2(input)

}

func part1(input string) {
	result := 1
	root := &Node{x: 0, y: 0}
	nodes := []Node{}
	nodes = append(nodes, *root)
	for _, letter := range input {
		isNew := true
		newNode := &Node{}
		if letter == '^' {
			newNode.y = root.y + 1
			newNode.x = root.x
			for _, node := range nodes {
				if newNode.x == node.x && newNode.y == node.y {
					isNew = false
					break
				}
			}
			if isNew {
				nodes = append(nodes, *newNode)
			}
		} else if letter == 'v' {
			newNode.y = root.y - 1
			newNode.x = root.x
			for _, node := range nodes {
				if newNode.x == node.x && newNode.y == node.y {
					isNew = false
					break
				}
			}
			if isNew {
				nodes = append(nodes, *newNode)
			}
		} else if letter == '>' {
			newNode.y = root.y
			newNode.x = root.x + 1
			for _, node := range nodes {
				if newNode.x == node.x && newNode.y == node.y {
					isNew = false
					break
				}
			}
			if isNew {
				nodes = append(nodes, *newNode)
			}
		} else if letter == '<' {
			newNode.y = root.y
			newNode.x = root.x - 1
			for _, node := range nodes {
				if newNode.x == node.x && newNode.y == node.y {
					isNew = false
					break
				}
			}
			if isNew {
				nodes = append(nodes, *newNode)
			}
		}
		root = newNode
	}
	result = len(nodes)
	fmt.Printf("Part 1: %d \n", result)
}

func part2(input string) {
	result := 1
	santa := &Node{x: 0, y: 0}
	root := &Node{x: 0, y: 0}
	nodes := []Node{}
	nodes = append(nodes, *santa)
	for i, letter := range input {
		if i%2 == 0 || i == 0 {
			isNew := true
			newNode := &Node{}
			if letter == '^' {
				newNode.y = santa.y + 1
				newNode.x = santa.x
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == 'v' {
				newNode.y = santa.y - 1
				newNode.x = santa.x
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == '>' {
				newNode.y = santa.y
				newNode.x = santa.x + 1
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == '<' {
				newNode.y = santa.y
				newNode.x = santa.x - 1
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			}
			santa = newNode
		} else {
			isNew := true
			newNode := &Node{}
			if letter == '^' {
				newNode.y = root.y + 1
				newNode.x = root.x
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == 'v' {
				newNode.y = root.y - 1
				newNode.x = root.x
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == '>' {
				newNode.y = root.y
				newNode.x = root.x + 1
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			} else if letter == '<' {
				newNode.y = root.y
				newNode.x = root.x - 1
				for _, node := range nodes {
					if newNode.x == node.x && newNode.y == node.y {
						isNew = false
						break
					}
				}
				if isNew {
					nodes = append(nodes, *newNode)
				}
			}
			root = newNode
		}
	}
	result = len(nodes)

	fmt.Printf("Part 2: %d \n", result)
}
