package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	/*
		暴力测试单调栈+滑动窗口是正确的 由于复杂度太高 非常耗时 queryTimes 设置小一点
	*/
	testTimes := 100
	arrLength := 2000
	queryTimes := 1000
	for i := 0; i < testTimes; i++ {
		fmt.Println("this is ", i, " times test")
		arr := generateRandomArray(arrLength)
		query := generate2DArray(queryTimes)
		ans1 := maxSubArrSum1(arrLength, queryTimes, arr, query)
		right := maxSubArrSumRight(arrLength, queryTimes, arr, query)
		if !isSame(ans1, right) {
			fmt.Println("Oops!!!!!")
			break
		}
	}
}

func TestSlidingWindowTime(t *testing.T) {
	/*
		单调栈+滑动窗口 平均耗时 6s 超时
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
		ans1 := maxSubArrSum1(arrLength, queryTimes, arr, query)
		if ans1 == nil {
			break
		}
		// 结束计时
		elapsed := time.Since(start)
		fmt.Printf("Execution time: %s\n", elapsed)
	}
}

/************************* 单调栈+滑动窗口 *************************/
// maxSubArrSum1 很可惜 时间复杂度太高 超时
func maxSubArrSum1(n, q int, arr []int, query [][]int) []int {
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		l := query[i][0]
		r := query[i][1]
		res := math.MinInt32
		var dq []int
		/*
			以 index 位置为结尾的子数组最大的累加和 这个角度思考问题
			注意进出单调栈的是前缀和
				进的时候是子数组长度刚好符合最短预期 那么进入的前缀和就是之前的一位 所以是 j-l
				同理 出的时候是子数组长度正好不符合预期的最长长度 所以是长度加一 要去掉 所以是 j-r-1
		*/
		for j := 1; j <= n; j++ {
			// 使用单调栈+滑动窗口
			// 进入单调栈
			if j-l >= 0 {
				idx := j - l
				for len(dq) > 0 && preSum[idx] <= preSum[dq[len(dq)-1]] {
					dq = dq[:len(dq)-1]
				}
				dq = append(dq, idx)
			}
			// 出单调栈
			if j-r-1 >= 0 {
				idx := j - r - 1
				if len(dq) > 0 && dq[0] == idx {
					dq = dq[1:]
				}
			}
			if len(dq) > 0 {
				res = max(res, preSum[j]-preSum[dq[0]])
			}
		}
		ans[i] = res
	}
	return ans
}
