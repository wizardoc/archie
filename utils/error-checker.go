package utils

import (
	"log"
)

func Check(err error, handler ...func()) {
	if err != nil {
		log.Fatal(err)
	}

	if len(handler) != 0 {
		handler[0]()
	}
}
