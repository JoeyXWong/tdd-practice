package string_calculator

import "strconv"

type StringCalculator interface {
	Add(inputOne string, inputTwo string) int
}

type stringCalculator struct {
}

func NewStringCalculator() StringCalculator {
	return &stringCalculator{}
}

func (this *stringCalculator) Add(inputOne, inputTwo string) int {
	sum := 0
	sum += this.checkIfEmpty(inputOne, inputTwo)
	sum += this.checkIfEmpty(inputTwo, inputOne)
	return sum
}

func (this *stringCalculator) checkIfEmpty(inputOne string, inputTwo string) int {
	if inputOne == "" {
		val, _ := strconv.ParseInt(inputTwo, 10, 32)
		return int(val)
	}
	return 0
}
