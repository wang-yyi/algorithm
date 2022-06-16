package leetcode

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"02", args{"a", "a"}, "a"},
		{"03", args{"a", "aa"}, ""},
		{"03", args{"你好世界你少点喝", "世界点"}, "世界你少点"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
