package day12

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func DFS(
	adj map[string][]string,
	current string,
	end string,
	path []string,
	paths *[][]string,
	visited map[string]int,
	allowance int,
) {
	visited[current]++
	path = append(path, current)
	if current == end {
		*paths = append(*paths, path)
		return
	}

	for _, neighbor := range adj[current] {
		if isLower(neighbor) && visited[neighbor] >= 1 && allowance < 1 {
			continue
		}
		if neighbor == "start" {
			continue
		}
		if isLower(neighbor) && visited[neighbor] == 1 {
			DFS(adj, neighbor, end, path, paths, visited, allowance-1)
		} else {
			DFS(adj, neighbor, end, path, paths, visited, allowance)
		}
		visited[neighbor]--
	}
}

func FindPaths(adj map[string][]string, start, end string) [][]string {
	var paths [][]string
	visited := make(map[string]int)
	DFS(adj, start, end, []string{}, &paths, visited, 0)
	return paths
}

func FindPathsCanVisitTwice(adj map[string][]string, start, end string) [][]string {
	var paths [][]string
	visited := make(map[string]int)
	DFS(adj, start, end, []string{}, &paths, visited, 1)
	return paths
}

func ParseInput(input io.ReadCloser) map[string][]string {
	adjacencyList := make(map[string][]string)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		from, to := parts[0], parts[1]
		adjacencyList[from] = append(adjacencyList[from], to)
		adjacencyList[to] = append(adjacencyList[to], from)
	}
	return adjacencyList
}

func isLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !unicode.IsLower(rune(s[i])) {
			return false
		}
	}
	return true
}
