package _17_letter_combinations_of_a_phone_number

var digitToLetters = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxzy",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var ans []string
	var cur = ""
	backtrace(digits, 0, cur, &ans)
	return ans
}

func backtrace(digits string, start int, cur string, ans *[]string) {
	if start == len(digits) {
		*ans = append(*ans, cur)
		return
	}

	digit := string(digits[start])
	letters := digitToLetters[digit]
	for i := 0; i < len(letters); i++ {
		backtrace(digits, start+1, cur+string(letters[i]), ans)
	}

}
