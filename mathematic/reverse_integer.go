package mathematic

import (
	"math"
)

//整数反转，中文版LeetCode，第7题

/*
解题思路：
算法思路：
  - 对当前数取对 10 的余数，再一项项填入res尾部，即可完成 int 翻转。
  - 边界情况处理： int 取值范围为 [−2^31,2^31−1] ，如果翻转数字溢出，则立即返回 0 。
    Python： 存储数字理论上是无限长度，因此每次计算完后判断res与of的大小关系即可；
*/
func reverse(x int) (rev int) {
	for x != 0 {
		if rev > math.MaxInt32/10 || x < math.MinInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
