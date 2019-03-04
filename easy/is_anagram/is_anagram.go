package is_anagram

type (
	IsAnagram interface {
		IsAnagram(input string, target string) bool
	}

	isAnagram struct {
	}
)

func NewIsAnagram() IsAnagram {
	return &isAnagram{}
}

func (this *isAnagram) IsAnagram(input string, target string) bool {
	if this.areInputsValid(input, target) {
		return false
	}
	anagram := this.mapAnagrams(input)
	for _, value := range target {
		if this.isLetterAnAnagram(anagram, value) {
			return false
		}
		anagram[value-'a']--
	}
	return true
}

func (this *isAnagram) isLetterAnAnagram(anagram map[int32]int, value int32) bool {
	return anagram[value-'a'] < 1
}

func (this *isAnagram) mapAnagrams(input string) map[int32]int {
	anagram := make(map[int32]int, len(input))
	for _, value := range input {
		anagram[value-'a']++
	}
	return anagram
}

func (this *isAnagram) areInputsValid(input string, target string) bool {
	return len(input) != len(target)
}
