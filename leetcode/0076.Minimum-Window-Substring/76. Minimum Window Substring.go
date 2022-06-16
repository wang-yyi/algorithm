package leetcode

import (
	"math"
)

func minWindow(s string, t string) string {
	s1, t1 := []rune(s), []rune(t)
	ori, cnt := map[rune]int{}, map[rune]int{}
	for _, c := range t1 {
		ori[c]++
	}

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	le := math.MaxInt32
	ansL, ansR := -1, -1
	l := 0
	for r, c := range s1 {
		if _, ok := ori[c]; ok {
			cnt[c]++
		}

		for l <= r && check() {
			if r-l+1 < le {
				le = r - l + 1
				ansL, ansR = l, r+1
			}

			if _, ok := ori[s1[l]]; ok {
				cnt[s1[l]]--
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}
	return string(s1[ansL:ansR])
}
