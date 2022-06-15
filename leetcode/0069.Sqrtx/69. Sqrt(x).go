package leetcode

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}

	}
	return ans
}
