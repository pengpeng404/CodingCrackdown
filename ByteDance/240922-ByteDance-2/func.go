package main

import (
	"math"
	"math/rand"
)

/************************* 暴力解 *************************/
// 这个暴力解是对的
func maxSubArrSumRight(n, q int, arr []int, query [][]int) []int {
	ans := make([]int, q)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}
	for i := 0; i < q; i++ {
		l := query[i][0]
		r := query[i][1]
		res := math.MinInt32
		for length := l; length <= r; length++ {
			for index := length; index <= n; index++ {
				res = max(res, preSum[index]-preSum[index-length])
			}
		}
		ans[i] = res
	}
	return ans
}

/************************* 辅助函数 *************************/
// 生成长度为n的随机数组，元素值范围为 -10^9 <= a_i <= 10^9
func generateRandomArray(n int) []int {
	// 定义数组
	arr := make([]int, n)
	// 填充数组，元素范围为 -10^9 <= a_i <= 10^9
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(2*1000000000) - 1000000000
	}
	return arr
}

// 生成长度为n的二维数组，每个子数组长度为2，且arr[i][0] <= arr[i][1]，数值范围为 1 到 2000
func generate2DArray(n int) [][]int {
	// 定义二维数组
	arr := make([][]int, n)

	// 填充二维数组
	for i := 0; i < n; i++ {
		a := rand.Intn(2000) + 1
		b := rand.Intn(2000) + 1

		// 确保 arr[i][0] <= arr[i][1]
		if a > b {
			a, b = b, a
		}

		arr[i] = []int{a, b}
	}

	return arr
}

// 判断两个切片是否相同
func isSame(a, b []int) bool {
	// 如果长度不同，直接返回false
	if len(a) != len(b) {
		return false
	}

	// 逐个元素比较
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
