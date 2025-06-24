package _17_letter_combinations_of_a_phone_number

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		digits string
		want   []string
	}{
		{
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			want:   []string{},
		},
		{
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			digits: "234",
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.digits, func(t *testing.T) {
			got := letterCombinations(tc.digits)

			// 排序后比较，忽略结果顺序
			sort.Strings(got)
			sort.Strings(tc.want)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("letterCombinations(%q) = %v, want %v", tc.digits, got, tc.want)
			}
		})
	}
}
