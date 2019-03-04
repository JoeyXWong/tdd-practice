package strstr

import "strings"

type StrStr interface {
	IndexOf(input string, target string) int
}

type strStr struct {
}

func NewStrStr() StrStr {
	return strStr{}
}

func (strStr) IndexOf(input string, target string) int {
	if len(strings.Trim(target, "")) == 0 {
		return 0
	}

	for i := range input {
		if len(input)-i < len(target) {
			return -1
		}
		for j := range target {
			if input[i+j] != target[j] {
				break
			}
			if j == len(target)-1 {
				return i
			}
		}
	}

	return -1
}
