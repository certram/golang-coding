package slice

import (
	"sort"
)

//合并区间，中文版LeetCode，第56题

/*
解题思路：
算法思路：
  - 先对区间slice按照第一位的数字进行升序排序（避免了，[[3,4],[2,8]]）这种情况，导致第三步还要多判断一次。
  - 把第一个区间添加进返回的结果list中，利用指针（len-1）指向这个list的末尾区间，
  - 第三步，判断这个末尾区间，判断这个区间的右部是否大于下一个区间的左部。如果结果list尾区间的右部小于下一个区间的左部，则2个区间没有重叠，把下一个区间添加到结果集中。如果区间有重叠，则找出结果集的尾部区间的右部和下一个要判断的区间的右部谁比较大，取最大的，完成合并。
  - 重复第三步；
  - 最后返回结果集

  运用到的包，就是sort.Slice([],func())
*/
//第一种写法
func MergeOne(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for _, e := range intervals[1:] {
		if ans[len(ans)-1][1] < e[0] {
			ans = append(ans, e)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], e[1])
		}
	}
	return
}

func MergeTwo(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = max(intervals[i][0], end)
		}
	}
	ans = append(ans, []int{start, end})
	return ans

}

// 返回两个整数较大的那个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
