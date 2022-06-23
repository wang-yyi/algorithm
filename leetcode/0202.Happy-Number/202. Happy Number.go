package leetcode

func isHappy(n int) bool {
	record := make(map[int]bool)
	getSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}

	for n != 1 && !record[n] {
		n, record[n] = getSum(n), true
	}
	return n == 1
}
