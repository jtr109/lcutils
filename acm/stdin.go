// Mocking stdin is inspired from https://stackoverflow.com/a/64518829/6522746

package acm

import (
	"os"
)

type MockedStdin struct {
	pipeReader *os.File
	pipeWriter *os.File
	oldStdin   *os.File
}

func MockStdin() (*MockedStdin, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	ms := &MockedStdin{
		pipeReader: r,
		pipeWriter: w,
		oldStdin:   os.Stdin,
	}
	os.Stdin = ms.pipeReader
	return ms, nil
}

func (ms *MockedStdin) Write(input string) error {
	content := []byte(input)
	if _, err := ms.pipeWriter.Write(content); err != nil {
		return err
	}
	if err := ms.pipeWriter.Close(); err != nil {
		return err
	}
	return nil
}

func (ms *MockedStdin) Close() (err error) {
	os.Stdin = ms.oldStdin // restore operate system standard input
	if err = ms.pipeReader.Close(); err != nil {
		return
	}
	return
}
