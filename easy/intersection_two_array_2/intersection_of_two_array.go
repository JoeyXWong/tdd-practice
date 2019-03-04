package intersection_two_array_2

type IntersectionOfTwoArray interface {
	GetIntersect(arrOne []int, arrTwo []int) []int
}

type intersectionOfTwoArray struct {
}

func NewIntersectionOfTwoArray() IntersectionOfTwoArray {
	return intersectionOfTwoArray{}
}

func (this intersectionOfTwoArray) GetIntersect(arrOne []int, arrTwo []int) []int {
	return this.computeIntersection(arrTwo, this.mapValuesOfArray(arrOne))
}

func (this intersectionOfTwoArray) computeIntersection(arrTwo []int, intersect map[int]int) []int {
	var result []int
	for _, value := range arrTwo {
		if intersect[value] > 0 {
			result = append(result, value)
		}
		intersect[value]--
	}
	return result
}

func (this intersectionOfTwoArray) mapValuesOfArray(arrOne []int) map[int]int{
	intersect := make(map[int]int)
	for _, value := range arrOne {
		intersect[value]++
	}
	return intersect
}
