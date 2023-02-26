package string

/*
 * @Author: certram
 * @Date: 2023-02-23 15:02:24
 * @LastEditors: certram
 * @LastEditTime: 2023-02-24 00:42:47
 */

//LeetCode第三题
/**
 * @funcname: lengthOfLongestSubstring
 * @description:
 * @param {string} s
 * @return {int}，返回最长子串的length
 */
func LengthOfLongestSubstring1(s string) int {
	//res记录最大的length，也是返回的结果
	p, res := -1, 0
	//声明一个临时map结构来标记一个字符是否出现过了。key是字符,value是代表是否出现过，可以是bool类型，我这里设置规则：value为0代表没出现过，为1代表出现过
	tmp := map[byte]int{}
	n := len(s)
	//i的意义：代表以字符串s的第i位开始的最长的子串，所以每一轮for循环都要删除s[i-1]
	// i为0，代表以s[0]开头的子串，
	// i为1，代表以s[1]开头的子串，
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(tmp, s[i-1])
		}
		//tmp[s[p+1]] == 0，这句代表了s[p+1]这个字符在map没有出现过
		for p+1 < n && tmp[s[p+1]] == 0 {
			tmp[s[p+1]]++ //这一句意义：添加标记，字符s[p+1]出现过了
			p++
		}
		//比较之前最长子串的length与当前子串的length，哪一个子串的length更大，取更大的那个length
		//经历for循环之后，p+1到达的位置：必然是以i开头的子串的第一个重复的字符的index,p就是这个子串的最后一个字符。
		//举例子，i在0位置，字符串s="pwdecccds",当p+1=5时，map里面c字符出现过了，不符合for循环条件，退出for。以0位置开始的不重复子串长度是：p-i+1
		res = max(res, p-i+1)
	}
	return res
}

// 解法2:位图
func LengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool

	res, l, r := 0, 0, 0
	for l < len(s) {
		if r < len(s) && bitSet[s[r]] == false {
			bitSet[s[r]] = true
			r++
		} else {
			bitSet[s[l]] = false
			l++

		}
		res = max(res, r-l)
	}
	return res
}

/**
 * @funcname:max
 * @description: 取两个int类型的数更大的那一个
 * @param {int} x
 * @param {int} y
 * @return {int} 返回两个数中更大的那个数
 */
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
