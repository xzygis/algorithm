package _2129_capitalize_the_title

import "strings"

// Easy
func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	var res []string
	for _, word := range words {
		if len(word) > 2 {
			res = append(res, strings.ToUpper(string(word[0]))+strings.ToLower(string(word[1:])))
		} else {
			res = append(res, strings.ToLower(word))
		}
	}

	return strings.Join(res, " ")
}
