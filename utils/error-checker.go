package utils

import (
	"fmt"
	"os"
)

func Check(err error, handler ...func()) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		if len(handler) != 0 {
			handler[0]()
		}
	}
}
