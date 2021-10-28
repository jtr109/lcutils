package acm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInput(t *testing.T) {
	input := "-10\n20\n5\n10\n3\n"
	mockedStdin, _ := MockStdin()
	defer mockedStdin.Close()
	mockedStdin.Write(input)

	assert.Equal(t, []int{-10, 20, 5, 10, 3}, readStdin())
}

func readStdin() []int {
	result := []int{}
	for {
		var a int
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			result = append(result, a)
		}
	}
	return result
}
