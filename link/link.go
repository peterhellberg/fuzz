// +build gofuzz

package link

import "github.com/peterhellberg/link"

func Fuzz(data []byte) int {
	if g := link.Parse(string(data)); g != nil {
		if len(g) > 0 {
			return 1
		}

		return 0
	}

	return -1
}
