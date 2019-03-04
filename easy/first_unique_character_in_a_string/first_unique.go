package first_unique_character_in_a_string

const NotFound = -1

type (
	FirstUniqueCharacter interface {
		FirstUnique(input string) int
	}

	firstUniqueCharacter struct {
	}
)

func NewFirstUniqueCharacter() FirstUniqueCharacter {
	return &firstUniqueCharacter{}
}

func (this *firstUniqueCharacter) FirstUnique(input string) int {

	unique := this.populateUniqueMapping(input)
	for i, value := range input {
		if this.isUnique(unique, value) {
			return i
		}
	}
	return NotFound
}

func (this *firstUniqueCharacter) isUnique(unique map[int32]int, value int32) bool {
	return unique[value-'a'] == 1
}

func (this *firstUniqueCharacter) populateUniqueMapping(input string) map[int32]int {
	unique := make(map[int32]int)
	for _, value := range input {
		unique[value-'a']++
	}
	return unique
}
