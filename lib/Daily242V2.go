package lib

// https://leetcode-cn.com/problems/valid-anagram/
// solution: https://leetcode-cn.com/problems/valid-anagram/solution/you-xiao-de-zi-mu-yi-wei-ci-by-leetcode-solution/

import (
	"fmt"
	"sort"
)

func toCharList(s string) []int {
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, int(s[i]))
	}
	return res
}

func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sList := toCharList(s)
	tList := toCharList(t)
	sort.Ints(sList)
	sort.Ints(tList)
	for i := 0; i < len(sList); i++ {
		if sList[i] != tList[i] {
			return false
		}
	}
	return true
}

func Run242V2() {
	res := isAnagramV2("sssddd", "sdsdsd")
	fmt.Println(res)
}
