// inspire from https://stackoverflow.com/a/64518829/6522746

package acm

import (
	"os"
)

type MockedStdin struct {
	pipeReader *os.File
	pipeWriter *os.File
	oldStdin   *os.File
}

func MockStdin1() (*MockedStdin, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	mockedStdin := &MockedStdin{
		pipeReader: r,
		pipeWriter: w,
		oldStdin:   os.Stdin,
	}
	return mockedStdin, nil
}

func (ms *MockedStdin) Write(input string) error {
	content := []byte(input)
	if _, err := ms.pipeWriter.Write(content); err != nil {
		return err
	}
	ms.pipeWriter.Close()
	os.Stdin = ms.pipeReader
	return nil
}

func (ms *MockedStdin) Close() (err error) {
	os.Stdin = ms.oldStdin // restore operate system standard input
	if err = ms.pipeReader.Close(); err != nil {
		return
	}
	return
}

func MockStdin(input string) (oldStdin *os.File, err error) {
	content := []byte(input)
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	if _, err = w.Write(content); err != nil {
		return nil, err
	}
	w.Close()

	oldStdin = os.Stdin
	os.Stdin = r
	return
}

func RestoreStdin(oldStdin *os.File) {
	os.Stdin = oldStdin // restore operate system standard input
}
