package lib

// https://leetcode-cn.com/problems/valid-anagram/

import "fmt"

func calcCharCount(s string, ) (m map[rune]uint) {
	m = make(map[rune]uint)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if _, ok := m[r]; !ok {
			m[r] = uint(0)
		}
		m[r] += 1
	}
	return
}

func isAnagram(s string, t string) bool {
	var (
		sMap = calcCharCount(s)
		tMap = calcCharCount(t)
	)
	if len(sMap) != len(tMap) {
		return false
	}
	for key, valS := range sMap {
		valT, ok := tMap[key]
		if !ok || valS != valT {
			return false
		}
	}
	return true
}

func Run242() {
	res := isAnagram("sssddd", "sdsdsd")
	fmt.Println(res)
}
