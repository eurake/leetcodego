package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for n := i + 1; n < len(nums); n++ {
			if nums[i]+nums[n] == target {
				return []int{i, n}
			}
		}
	}
	return []int{}

}

func main() {
	nums := []int{2, 7, 11, 15}

	indexs := twoSum(nums, 9)

	fmt.Println(indexs)
}
