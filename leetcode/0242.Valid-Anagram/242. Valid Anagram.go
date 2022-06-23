package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	exists := make(map[rune]int)
	for _, v := range s {
		exists[v]++
	}

	for _, v := range t {
		exists[v]--
	}

	for _, v := range exists {
		if v != 0 {
			return false
		}
	}
	return true
}
