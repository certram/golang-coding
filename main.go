/*
 * @Author: certram
 * @Date: 2023-02-23 15:25:19
 * @LastEditors: certram
 * @LastEditTime: 2023-02-24 00:47:28
 */
package main

import (
	"fmt"

	s1 "github.com/certram/golang-coding/string"
)

func main() {
	s := "fedywewd"
	n := s1.LengthOfLongestSubstring2(s)

	fmt.Println(n)
}
