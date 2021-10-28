# ACM Mode Utils

## Usage

### Mock Stdin

```go
package example

import "github.com/jtr109/lcutils/acm"

func main() {
	input := "-10\n20\n5\n10\n3\n"
	mockedStdin, _ := acm.MockStdin()
	defer mockedStdin.Close()
	mockedStdin.Write(input)

	output := readStdin() // []int{-10, 20, 5, 10, 3}
}
```

### Mock Stdout

```go
import "github.com/jtr109/lcutils/acm"

func main() {
	mockedStdout, _ := acm.MockStdout()
	defer mockedStdout.Close()

	content := "Hello, playground"
	fmt.Print(content)

	out, _ := mockedStdout.Read() // "Hello, playground"
}
```
