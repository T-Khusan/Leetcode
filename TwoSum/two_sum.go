package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}

func main() {
	n := twoSum([]int{2,7,11,15}, 9)
	fmt.Println(n)
}
