package template

import (
	"bufio"
	"fmt"
	"os"
)

func dfs2(graph map[int][]int, visited []bool, originNode int) {
	if visited[originNode] {
		return
	}
	visited[originNode] = true
	path, ok := graph[originNode]
	if !ok {
		return
	}
	for _, target := range path {
		dfs2(graph, visited, target)
	}
}

func bfs2(graph map[int][]int, visited []bool, originNode int) {
	// if visited[originNode] {
	// 	return
	// }
	visited[originNode] = true
	queue := make([]int, 0)
	queue = append(queue, originNode)
	for len(queue) != 0 {
		origin := queue[0]
		//如果在这里标记，在节点入队时有可能回一个节点多次入队，因为入队时节点还没有标记
		// visited[origin] = true
		queue = queue[1:]

		path, ok := graph[origin]
		if !ok {
			continue
		}
		for _, value := range path {
			if !visited[value] {
				queue = append(queue, value)
				//在节点入队时标记效果更好能剪枝
				visited[value] = true
			}
		}
	}
}

func TestBfs2() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

	graph := make(map[int][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		var s, t int
		fmt.Sscanf(scanner.Text(), "%d %d", &s, &t)
		graph[s] = append(graph[s], t)
	}

	visited := make([]bool, n+1)

	bfs2(graph, visited, 1)

	for i := 1; i <= n; i++ {
		if !visited[i] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(1)

}
