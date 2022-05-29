package leetcode

func Rotate(nums []int, k int) {
	if k > 0 {
		last := nums[len(nums) - 1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i - 1]
		}
		nums[0] = last
		Rotate(nums, (k - 1))
	}
}
