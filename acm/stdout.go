// Mocking stdout is inspired from https://stackoverflow.com/a/29339052/6522746

package acm

import (
	"io/ioutil"
	"os"
)

type MockedStdout struct {
	pipeReader *os.File
	pipeWriter *os.File
	oldStdout  *os.File
}

func MockStdout() (*MockedStdout, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	ms := &MockedStdout{
		pipeReader: r,
		pipeWriter: w,
		oldStdout:  os.Stdout,
	}
	os.Stdout = ms.pipeWriter
	return ms, nil
}

func (ms *MockedStdout) Read() (string, error) {
	ms.pipeWriter.Close()
	out, _ := ioutil.ReadAll(ms.pipeReader)
	return string(out), nil
}

func (ms *MockedStdout) Close() (err error) {
	os.Stdout = ms.oldStdout // restore original Stdout
	err = ms.pipeReader.Close()
	return
}
