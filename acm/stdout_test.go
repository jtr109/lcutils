package acm

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureStdout(t *testing.T) {
	// pipeReader, pipeWriter, oldStdout, _ := MockStdout1()
	mockedStdout, err := MockStdout()
	if err != nil {
		log.Fatal(err)
	}
	defer mockedStdout.Close()

	content := "Hello, playground"
	fmt.Print(content) // this gets captured

	out, _ := mockedStdout.Read()

	assert.Equal(t, content, out)
}
