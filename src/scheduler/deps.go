package scheduler

import "fmt"

// DetectCycles returns a list of cycles in the dependency graph.
func DetectCycles(graph map[string][]string) [][]string {
	var cycles [][]string
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	for file := range graph {
		if !visited[file] {
			pathStack := []string{}
			if detectCycle(file, graph, visited, recStack, pathStack) {
				cycles = append(cycles, pathStack)
			}
		}
	}
	return cycles
}

func detectCycle(path string, deps map[string][]string, visited, recStack map[string]bool, pathStack []string) bool {
	visited[path] = true
	recStack[path] = true
	pathStack = append(pathStack, path)
	for _, dep := range deps[path] {
		if !visited[dep] {
			if detectCycle(dep, deps, visited, recStack, pathStack) {
				return true
			}
		} else if recStack[dep] {
			return true
		}
	}
	recStack[path] = false
	return false
}

// TopoSort returns a topologically sorted list of files based on the dependency graph.
func TopoSort(graph map[string][]string) ([]string, error) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	var result []string
	for file := range graph {
		if !visited[file] {
			if !topoSortDFS(file, graph, visited, recStack, &result) {
				return nil, ErrCycleDetected
			}
		}
	}
	return result, nil
}

var ErrCycleDetected = fmt.Errorf("circular dependency detected")

func topoSortDFS(file string, graph map[string][]string, visited, recStack map[string]bool, result *[]string) bool {
	visited[file] = true
	recStack[file] = true
	for _, dep := range graph[file] {
		if !visited[dep] {
			if !topoSortDFS(dep, graph, visited, recStack, result) {
				return false
			}
		} else if recStack[dep] {
			return false
		}
	}
	recStack[file] = false
	*result = append(*result, file)
	return true
}

// NextBatch returns all files whose dependencies are completed and not yet processed.
func NextBatch(graph map[string][]string, completed map[string]bool, scheduled map[string]bool) []string {
	var batch []string
	for file, deps := range graph {
		if completed[file] || scheduled[file] {
			continue
		}
		ready := true
		for _, dep := range deps {
			if !completed[dep] {
				ready = false
				break
			}
		}
		if ready {
			batch = append(batch, file)
		}
	}
	return batch
}