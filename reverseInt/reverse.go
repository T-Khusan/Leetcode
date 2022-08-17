package main

import (
	"fmt"
)

func reverse(x int) int {
	result, sign := 0, 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x = x / 10
	}

	var checkInt int = int(int32(result))

	if checkInt != result {
		return 0
	}

	return result * sign
}

func main() {
	fmt.Println(reverse(-100))
	fmt.Println(reverse(-1004))
	fmt.Println(reverse(131415))
	fmt.Println(reverse(1357))
}

// github
/*
package reverse_integer

func reverse(x int) int {

	if (x > 1<<31-1) || (x < -1<<31) {
		return 0
	}

	var res = 0
	for x != 0 {
		last := x % 10
		res = (res * 10) + last
		x = x / 10
	}

	if (res > 1<<31-1) || (res < -1<<31) {
		return 0
	}

	return res
}
*/

// https://github.com/rorua/leetcode_go
