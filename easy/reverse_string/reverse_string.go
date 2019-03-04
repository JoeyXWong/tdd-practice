package reverse_string

type (
	ReverseString interface {
		Reverse(input []byte) []byte
	}

	reverseString struct {
	}
)

func NewReverseString() ReverseString {
	return reverseString{}
}

func (reverseString) Reverse(input []byte) []byte {
	end := len(input) - 1
	start := 0
	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}
	return input
}
