package leetcode

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)

	for i := range nums {
		if _, ok := dict[nums[i]]; ok {
			return true
		}
		dict[nums[i]] = nums[i]
	}
	return false
}
