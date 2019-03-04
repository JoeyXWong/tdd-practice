package rotate_image

type (
	RotateImage interface {
		Rotate(matrix [][]int) [][]int
	}

	rotateImage struct {
	}
)

func NewRotateImage() RotateImage {
	return &rotateImage{}
}

func (this *rotateImage) Rotate(matrix [][]int) [][]int {
	this.transpose(matrix)
	this.swapXAxis(matrix)
	return matrix
}

func (this *rotateImage) swapXAxis(matrix [][]int) {
	for i := range matrix {
		start := 0
		end := len(matrix[i]) - 1
		for start < end {
			matrix[i][start], matrix[i][end] = matrix[i][end], matrix[i][start]
			start++
			end--
		}
	}
}

func (this *rotateImage) transpose(matrix [][]int) {
	for i := range matrix {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
