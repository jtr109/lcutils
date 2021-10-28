package acm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureStdout(t *testing.T) {
	mockedStdout, _ := MockStdout()
	defer mockedStdout.Close()

	content := "Hello, playground"
	fmt.Print(content) // this gets captured

	out, _ := mockedStdout.Read()

	assert.Equal(t, content, out)
}
