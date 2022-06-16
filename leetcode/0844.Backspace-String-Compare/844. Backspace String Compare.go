package leetcode

func backspaceCompare(s string, t string) bool {
	s1 := make([]rune, 0)
	for _, c := range s {
		if c == '#' {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			}
		} else {
			s1 = append(s1, c)
		}
	}
	t1 := make([]rune, 0)
	for _, c := range t {
		if c == '#' {
			if len(t1) > 0 {
				t1 = t1[:len(t1)-1]
			}
		} else {
			t1 = append(t1, c)
		}
	}
	return string(s1) == string(t1)
}
