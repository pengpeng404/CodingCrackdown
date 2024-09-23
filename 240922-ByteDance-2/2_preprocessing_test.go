package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestPreprocessing(t *testing.T) {
	/*
		暴力测试预处理数据是正确的
	*/
	testTimes := 100
	arrLength := 2000
	queryTimes := 1000
	for i := 0; i < testTimes; i++ {
		fmt.Println("this is ", i, " times test")
		arr := generateRandomArray(arrLength)
		query := generate2DArray(queryTimes)
		ans2 := maxSubArrSumPreprocessing(arrLength, queryTimes, arr, query)
		right := maxSubArrSumRight(arrLength, queryTimes, arr, query)
		if !isSame(ans2, right) {
			fmt.Println("Oops!!!!!")
			break
		}
	}
}

func TestPreprocessingTime(t *testing.T) {
	/*
		暴力预处理 平均耗时 250ms 我是傻逼
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
		ans2 := maxSubArrSumPreprocessing(arrLength, queryTimes, arr, query)
		if ans2 == nil {
			break
		}
		// 结束计时
		elapsed := time.Since(start)
		fmt.Printf("Execution time: %s\n", elapsed)
	}
}

/************************* 预处理 *************************/
func maxSubArrSumPreprocessing(n, q int, arr []int, query [][]int) []int {
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
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		l := query[i][0]
		r := query[i][1]
		res := math.MinInt32
		for j := l; j <= r; j++ {
			res = max(res, maxSumOfLen[j])
		}
		ans[i] = res
	}
	return ans
}
