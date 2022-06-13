package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"01", args{nums: []int{1, 3, 5, 6}, target: 5}, 2},
		{"02", args{nums: []int{1, 3, 5, 6}, target: 2}, 1},
		{"03", args{nums: []int{1, 3, 5, 6}, target: 7}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
