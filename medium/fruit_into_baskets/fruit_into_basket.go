package fruit_into_baskets

type FruitIntoBasket interface {
	GetFruits(fruits []int) int
}

type fruitIntoBasket struct {
}

func NewFruitIntoBasket() FruitIntoBasket {
	return &fruitIntoBasket{}
}

func (*fruitIntoBasket) GetFruits(fruits []int) int {
	basket := make(map[int]int)
	distinctFruits := 0
	fruitCount := 0

	start := 0
	end := 0

	for end < len(fruits) {
		if basket[end] == 0 {
			distinctFruits++
			basket[end]++
			if fruitCount < end - start + 1{
				fruitCount = end - start + 1
			}
		}


		if distinctFruits < 3 {
			end++
		} else {
			basket[start]--
			if basket[start] == 0{
				distinctFruits--
			}
			start++
		}
	}
	return fruitCount
}
