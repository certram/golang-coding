/*
 * @Author: certram
 * @Date: 2023-02-23 14:21:47
 * @LastEditors: certram
 * @LastEditTime: 2023-02-26 15:41:59
 */
package array

/*
 * @Author: certram
 * @Date: 2023-02-23 14:21:47
 * @LastEditors: certram
 * @LastEditTime: 2023-02-26 15:25:59
 */

// LeetCode 的第1题
/**
 * @name:
 * @description:
 * @param {[]int} nums
 * @param {int} target
 * @return {*}
 */
func twoSum(nums []int, target int) []int {
	//创建一个map来存储已经遍历过的nums的index,以及index对应的值。
	storageMap := map[int]int{}
	//遍历nums，看看map里面有没有key值符合条件：和为target,那么要到map里面找的是target-value
	for index, value := range nums {
		// 查查看storageMap里面是否存有值：target-value，如果有，返回这两个数的index
		if v, ok := storageMap[target-value]; ok { //这个v，ok中的v其实就是nums中的下标
			//找到了和为target的那个数了，返回两个值的下标
			return []int{v, index}
		}
		//没有找到的情况，把这个value,及对应的index加入到storageMap当中,storageMap的key是nums的值，value是值的下标index
		storageMap[value] = index

	}
	//遍历完nums中的值，都没有找到，那么不存在这样的两个数
	return nil
}
