/*
 * @Author: certram
 * @Date: 2023-02-26 12:38:24
 * @LastEditors: certram
 * @LastEditTime: 2023-02-26 15:52:51
 */
package array

import (
	"fmt"
	"math"
)

func ReverseInt(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	fmt.Println("ii")
	return
}
