package main

import (
	"fmt"
)

func main() {
	x := 2123456789008887779
	fmt.Printf("%d\n", reverse(x))
	fmt.Printf("%T\n", reverse(x))
}
func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	return tmp
}
