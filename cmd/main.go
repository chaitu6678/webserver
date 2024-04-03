package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// Parse the string array into a graph represented as an adjacency list.
func parseGraph(strArr []string, n int) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range strArr[n+1:] {
		nodes := strings.Split(edge, "-")
		if len(nodes) == 2 {
			graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
			// If it's an undirected graph, add the edge in the opposite direction too.
			graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
		}
	}
	return graph
}

// Find the shortest path using BFS.
func bfs(graph map[string][]string, start, end string) []string {
	if start == end {
		return []string{start}
	}

	queue := list.New()
	visited := make(map[string]bool)
	prev := make(map[string]string)

	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(string)
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				prev[neighbor] = current
				if neighbor == end {
					// End found, stop the search
					queue = list.New()
					break
				}
			}
		}
	}

	path := make([]string, 0)
	at := end
	for at != "" {
		path = append([]string{at}, path...)
		at = prev[at]
		if at == start {
			path = append([]string{start}, path...)
			return path
		}
	}

	// No path found
	return []string{"-1"}
}

// ShortestPath finds the shortest path in a graph.
func ShortestPath(strArr []string) string {
	n, _ := strconv.Atoi(strArr[0]) // The number of nodes N as an integer
	if n == 0 {
		return "-1"
	}

	graph := parseGraph(strArr, n)
	start := strArr[1]
	end := strArr[n]

	shortestPath := bfs(graph, start, end)
	if len(shortestPath) == 0 {
		return "-1"
	}
	return strings.Join(shortestPath, "-")
}

func main() {
	// Example input
	strArr := []string{"6", "A", "B", "C", "D", "E", "F", "A-B", "B-C", "C-D", "D-E", "E-F"}
	fmt.Println("Shortest Path:", ShortestPath(strArr))
}
