# fuzz

Randomized testing for my Go packages using [go-fuzz](https://github.com/dvyukov/go-fuzz)

## Example

```go
// +build gofuzz

package duration

import "github.com/peterhellberg/duration"

func Fuzz(data []byte) int {
	if _, err := duration.Parse(string(data)); err == nil {
		return 1
	}

	return -1
}
```
