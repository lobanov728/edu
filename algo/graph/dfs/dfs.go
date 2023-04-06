package main

import "fmt"

func main() {
	graph := map[int][]int{
		1:  {2, 5, 4, 8, 3},
		2:  {1, 5},
		3:  {1, 7},
		4:  {1, 8},
		5:  {9, 1, 2},
		7:  {3, 11},
		8:  {1, 4, 13, 12},
		9:  {5, 13},
		11: {7, 12},
		12: {8, 11, 15, 25},
		13: {8, 9},
		15: {31, 12},
		24: {31, 25},
		25: {24, 12},
		31: {15, 53},
		53: {31, 60, 69},
		60: {53, 73},
		69: {53, 73},
	}

	visited := []int{}
	fmt.Println(dfs(1, 73, graph, visited, 0))
	fmt.Println()
	fmt.Println(dfsIterative(1, 73, graph))
}

func dfs(startNode int, searchNode int, graph map[int][]int, visited []int, depth int64) bool {
	fmt.Println("startNode", startNode)
	// fmt.Println("visited1", visited)
	if startNode == searchNode {
		return true
	}

	if inArray(startNode, visited) {
		return false
	}

	visited = append(visited, startNode)

	for i := range graph[startNode] {
		// fmt.Println("graph[startNode][i]", graph[startNode][i])
		// fmt.Println("depth", depth)
		// fmt.Println("visited2", visited)
		if !inArray(graph[startNode][i], visited) {
			nDepth := depth + 1
			if dfs(graph[startNode][i], searchNode, graph, visited, nDepth) {
				return true
			}
		}
	}

	return false
}

func inArray(node int, visited []int) bool {
	for i := range visited {
		if node == visited[i] {
			return true
		}
	}

	return false
}

func dfsIterative(startNode int, searchNode int, graph map[int][]int) bool {
	stack := make([]int, 0, 100)
	stack = append(stack, startNode)
	visited := make([]int, 25)

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node == searchNode {
			fmt.Println("FOUNDED")
			return true
		}

		if !inArray(node, visited) {
			fmt.Println(node)
			visited = append(visited, node)
			adjacent := graph[node]
			stack = append(adjacent, stack...)
		}
	}
	fmt.Println("NOT FOUNDED")
	return false
}
