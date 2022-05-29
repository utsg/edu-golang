package leetcode

func removeDuplicates(nums []int) int {
	result := 0
	
	for i := 0; i < len(nums); i++ {
		if (nums[result] != nums[i]) {
			result++
			nums[result] = nums[i]
		}
	}
	return result + 1
}
