package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// parseGraph creates an adjacency list from the provided string array
func parseGraph(strArr []string, n int) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range strArr[1+n:] {
		parts := strings.Split(edge, "-")
		if len(parts) == 2 {
			graph[parts[0]] = append(graph[parts[0]], parts[1])
			graph[parts[1]] = append(graph[parts[1]], parts[0])
		}
	}
	return graph
}

// bfs performs a breadth-first search on the graph to find the shortest path
func bfs(graph map[string][]string, start, end string) []string {
	visited := make(map[string]bool)
	prev := make(map[string]string)
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(string)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = node
				queue.PushBack(neighbor)
				if neighbor == end {
					queue = list.New() // Found the end; clear the queue to break the loop
					break
				}
			}
		}
	}

	var path []string
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}

	if path[0] != start { // Check if a path was found
		return []string{}
	}

	return path
}

// ShortestPath finds the shortest path from the first to the last node in strArr
func ShortestPath(strArr []string) string {
	n, _ := strconv.Atoi(strArr[0]) // Number of nodes
	if n <= 0 {
		return "No valid nodes provided"
	}

	graph := parseGraph(strArr, n)
	start := strArr[1]
	end := strArr[n]

	path := bfs(graph, start, end)
	if len(path) == 0 {
		return "No path found"
	}

	return strings.Join(path, "-")
}

func main() {
	// Example input
	strArr := []string{"6", "A", "B", "C", "D", "E", "F", "A-B", "B-C", "C-D", "D-E", "E-F"}
	fmt.Println("Shortest Path:", ShortestPath(strArr))
}
