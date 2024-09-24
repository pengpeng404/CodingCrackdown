package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 3, 2}
	fmt.Println(maxBridges(arr1, arr2))
}

func maxBridges(arr1, arr2 []int) int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}
	//return process(arr1, arr2, len(arr1), len(arr2), 0, 0)
	m := len(arr1)
	n := len(arr2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			p1 := 0
			if arr1[i] == arr2[j] {
				p1 = 1 + dp[i+1][j+1]
			}
			dp[i][j] = max(p1, dp[i+1][j], dp[i][j+1])
		}
	}
	return dp[0][0]
}

// process 桥墩从 idx1 idx2 开始到最后可以最多组成多少个桥墩
func process(arr1, arr2 []int, m, n, idx1, idx2 int) int {
	if idx1 == m || idx2 == n {
		return 0
	}
	p1 := 0
	p2 := 0
	p3 := 0
	if arr1[idx1] == arr2[idx2] {
		p1 = 1 + process(arr1, arr2, m, n, idx1+1, idx2+1)
	} else {
		p2 = process(arr1, arr2, m, n, idx1, idx2+1)
		p3 = process(arr1, arr2, m, n, idx1+1, idx2)
	}
	return max(p1, p2, p3)
}
