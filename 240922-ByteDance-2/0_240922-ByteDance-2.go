package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	/*
		预处理数组+线段树 正确
	*/
	//testTimes := 100
	//arrLength := 2000
	//queryTimes := 1000
	//for i := 0; i < testTimes; i++ {
	//	fmt.Println("this is ", i, " times test")
	//	arr := generateRandomArray(arrLength)
	//	query := generate2DArray(queryTimes)
	//	ans2 := maxSubArrSum(arrLength, queryTimes, arr, query)
	//	right := maxSubArrSumRight(arrLength, queryTimes, arr, query)
	//	if !isSame(ans2, right) {
	//		fmt.Println(ans2)
	//		fmt.Println(right)
	//		fmt.Println("Oops!!!!!")
	//		break
	//	}
	//}

	/*
		测试时间 平均耗时 120ms
	*/
	testTimes := 100
	arrLength := 2000
	queryTimes := 1000000
	for i := 0; i < testTimes; i++ {
		fmt.Println("this is ", i, " times")
		arr := generateRandomArray(arrLength)
		query := generate2DArray(queryTimes)
		// 开始计时
		start := time.Now()
		ans1 := maxSubArrSum(arrLength, queryTimes, arr, query)
		if ans1 == nil {
			break
		}
		// 结束计时
		elapsed := time.Since(start)
		fmt.Printf("Execution time: %s\n", elapsed)
	}

}

/************************* 预处理数组+线段树 *************************/
func maxSubArrSum(n, q int, arr []int, query [][]int) []int {
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}
	maxSumOfLen := make([]int, n+1)
	for length := 1; length <= n; length++ {
		res := math.MinInt32
		for i := length; i <= n; i++ {
			res = max(res, preSum[i]-preSum[i-length])
		}
		maxSumOfLen[length] = res
	}
	sg := buildSegmentTree(maxSumOfLen, 1, n)
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		ans[i] = sg.query(query[i][0], query[i][1])
	}
	return ans
}

type segmentTreeNode struct {
	start, end int
	maxNum     int
	left       *segmentTreeNode
	right      *segmentTreeNode
}

func buildSegmentTree(arr []int, start, end int) *segmentTreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &segmentTreeNode{
			start:  start,
			end:    end,
			maxNum: arr[start],
		}
	}
	node := &segmentTreeNode{start: start, end: end}
	mid := (start + end) / 2
	node.left = buildSegmentTree(arr, start, mid)
	node.right = buildSegmentTree(arr, mid+1, end)
	node.maxNum = max(node.left.maxNum, node.right.maxNum)
	return node
}

func (sg *segmentTreeNode) query(l, r int) int {
	if sg == nil || l > sg.end || r < sg.start {
		return math.MinInt32
	}
	if l <= sg.start && r >= sg.end {
		return sg.maxNum
	}
	leftMax := sg.left.query(l, r)
	rightMax := sg.right.query(l, r)
	return max(leftMax, rightMax)
}
