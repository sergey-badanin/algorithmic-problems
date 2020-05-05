package groupanagrams

import (
	"container/list"
	"sort"
)

//Solution for the problem:
//Given an array of strings, group anagrams together
//
//In the solution calculate "hash" by ordering bytes of every string and group them accordingly
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string]*list.List)
	for _, v := range strs {
		sorted := sortBytes(v)
		anagram, ok := anagrams[sorted]
		if ok {
			anagram.PushBack(v)
		} else {
			elem := list.New()
			elem.PushBack(v)
			anagrams[sorted] = elem
		}
	}

	result := make([][]string, len(anagrams))
	i := 0
	for _, v := range anagrams {
		line := make([]string, v.Len())
		j := 0
		for e := v.Front(); e != nil; e = e.Next() {
			line[j] = e.Value.(string)
			j++
		}
		result[i] = line
		i++
	}

	return result
}

func sortBytes(s string) string {
	a := make([]byte, len(s))
	copy(a, s)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return string(a)
}
