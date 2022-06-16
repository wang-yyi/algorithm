package leetcode

func sortedSquares(nums []int) []int {
	n := len(nums)
	low, height, k := 0, n-1, n-1
	ans := make([]int, n)
	for low <= height {
		n1 := nums[low] * nums[low]
		n2 := nums[height] * nums[height]
		if n1 > n2 {
			ans[k] = n1
			low++
		} else {
			ans[k] = n2
			height--
		}
		k--
	}
	return ans
}
