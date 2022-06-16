package leetcode

func totalFruit(fruits []int) int {
	left, res, hashMap := 0, 0, map[int]int{}
	for i := 0; i < len(fruits); i++ {
		hashMap[fruits[i]]++
		for len(hashMap) > 2 {
			hashMap[fruits[left]]--
			if hashMap[fruits[left]] == 0 {
				delete(hashMap, fruits[left])
			}
			left++
		}

		num := i - left + 1
		if num > res {
			res = num
		}
	}
	return res
}
