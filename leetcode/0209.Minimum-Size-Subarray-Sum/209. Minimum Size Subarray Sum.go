package leetcode

func minSubArrayLen(target int, nums []int) int {
	left, num, res := 0, 0, len(nums)+1
	for right := 0; right < len(nums); right++ {
		num += nums[right]
		for num >= target {
			n := right - left + 1
			if n < res {
				res = n
			}
			num -= nums[left]
			left++
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}
