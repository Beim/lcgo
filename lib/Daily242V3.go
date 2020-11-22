package lib

// https://leetcode-cn.com/problems/valid-anagram/
// solution: https://leetcode-cn.com/problems/valid-anagram/solution/you-xiao-de-zi-mu-yi-wei-ci-by-leetcode-solution/

import (
	"fmt"
)

func isAnagramV3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i] - 'a'] += 1
		hash[t[i] - 'a'] -= 1
	}
	for i := 0; i < 26; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}

func Run242V3() {
	res := isAnagramV3("sssddd", "sdsdsd")
	fmt.Println(res)
}
