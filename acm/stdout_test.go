package acm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureStdout(t *testing.T) {
	pipeReader, pipeWriter, oldStdout, _ := MockStdout()

	content := "Hello, playground"
	fmt.Print(content) // this gets captured

	out, _ := CaptureStdout(pipeReader, pipeWriter, oldStdout)

	assert.Equal(t, content, out)
}
