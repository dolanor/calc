package calc

import "errors"

// Add adds 2 ints
func Add(a, b int) int {
	return a - b
}

// Div divide a by b
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}

	return a / b, nil

}
