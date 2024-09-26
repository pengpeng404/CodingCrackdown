package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 定义节点结构体
type node struct {
	val   int
	edges []int
}

// 深度优先遍历
func dfs(currentNode *node, depth, k int, nodes map[int]*node, visited map[int]bool, set map[int]int) {
	if depth > k {
		return
	}
	set[currentNode.val] = 1
	visited[currentNode.val] = true
	for _, neighbor := range currentNode.edges {
		if !visited[neighbor] {
			dfs(nodes[neighbor], depth+1, k, nodes, visited, set)
		}
	}
}

func main() {
	// 创建快速输入输出
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	// 读取第一行：n, m, k
	reader.Scan()
	firstLine := strings.Split(reader.Text(), " ")
	n, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])
	k, _ := strconv.Atoi(firstLine[2])

	// 构建图
	nodes := make(map[int]*node)
	for i := 0; i < n; i++ {
		reader.Scan()
		line := strings.Split(reader.Text(), " ")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])

		if _, ok := nodes[x]; !ok {
			nodes[x] = &node{val: x}
		}
		if _, ok := nodes[y]; !ok {
			nodes[y] = &node{val: y}
		}
		nodes[x].edges = append(nodes[x].edges, y)
		nodes[y].edges = append(nodes[y].edges, x)
	}

	// 使用集合 set 来记录在 K 跳内到达的节点
	set := make(map[int]int)
	visited := make(map[int]bool)

	// 从指定的用户 m 开始搜索
	dfs(nodes[m], 0, k, nodes, visited, set)

	// 输出影响力，减去自己
	_, _ = fmt.Fprintln(writer, len(set)-1)
}
