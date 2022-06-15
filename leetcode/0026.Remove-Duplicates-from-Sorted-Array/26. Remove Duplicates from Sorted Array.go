package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slowIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[slowIndex] = nums[i]
			slowIndex++
		}
	}
	return slowIndex
}
