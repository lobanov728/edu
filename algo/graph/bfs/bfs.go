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
		24: {31, 25},
		25: {24, 12},
		31: {15, 53},
		53: {31, 60, 69},
		60: {53, 73},
		69: {53, 73},
	}

	bfs(1, 73, graph)

}

func bfs(startNode int, searchNode int, graph map[int][]int) bool {
	visited := make([]int, 25)

	var searchQueue []int
	searchQueue = graph[startNode]
	for len(searchQueue) > 0 {
		node := searchQueue[0]
		searchQueue = searchQueue[1:]

		if !inArray(node, visited) {
			fmt.Println(node)
			if node == searchNode {
				fmt.Println("FOUNDED")
				return true
			}
			searchQueue = append(searchQueue, graph[node]...)
			visited = append(visited, node)
		}
	}

	fmt.Println("NOT FOUND")
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
