package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := record[v]; !ok {
			record[v] = struct{}{}
		}
	}

	ret := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := record[v]; ok {
			ret = append(ret, v)
			delete(record, v)
		}
	}
	return ret
}
