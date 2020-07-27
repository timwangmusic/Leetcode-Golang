package backtracking

import (
	"strings"
)

func letterCombinations(digits string) []string {
	letters := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := make([]string, 0)
	path := make([]string, 0)
	d := strings.Split(digits, "")
	backtrack(letters, path, d, 0, &res)
	return res
}

func backtrack(letters map[string][]string, path []string, digits []string, idx int, res *[]string) {
	if idx == len(digits) {
		tmp := strings.Join(path, "")
		if len(tmp) > 0 {
			*res = append(*res, tmp)
		}
		return
	}

	key := digits[idx]
	for _, letter := range letters[key] {
		path = append(path, letter)
		backtrack(letters, path, digits, idx+1, res)
		_, path = path[len(path)-1], path[:len(path)-1]
	}
}
