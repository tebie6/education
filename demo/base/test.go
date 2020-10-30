package main

import "fmt"

func main(){
	var nums = []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(nums)
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	nui := 0
	for i := 1; i < len(nums); i++ {
		if nums[nui] != nums[i] {
			nui++
			nums[nui] = nums[i]
		}
	}
	return nui + 1
}
