package lcof

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n
		}
		m[n] = true
	}
	return 0
}
