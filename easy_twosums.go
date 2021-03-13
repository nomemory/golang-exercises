package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	var cache = make(map[int]int)
	var search int
	for idx, el := range nums {
		search = target - el
		if searchIdx, ok := cache[search]; ok {
			var res []int = []int{idx, searchIdx}
			sort.Ints(res)
			return res
		}
		cache[el] = idx
	}
	return []int{0, 0}
}

func main() {
	var nums []int = []int{3, 2, 4}
	var target int = 6
	fmt.Println(twoSum(nums, target))
}
