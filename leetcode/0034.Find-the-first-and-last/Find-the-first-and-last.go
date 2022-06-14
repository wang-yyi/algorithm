package leetcode

func searchRange(nums []int, target int) []int {
	getRight := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		border := -2
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
				border = left
			}
		}
		return border
	}

	getLeft := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		border := -2
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] >= target {
				right = mid - 1
				border = right
			} else {
				left = mid + 1
			}
		}
		return border
	}

	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)

	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}

	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}

	return []int{-1, -1}
}