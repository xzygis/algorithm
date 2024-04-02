package groupanagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	group := map[string][]string{}
	for _, str := range strs {
		key := format(str)
		group[key] = append(group[key], str)
	}

	res := make([][]string, 0, len(group))
	for _, v := range group {
		res = append(res, v)
	}

	return res
}

func format(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}
