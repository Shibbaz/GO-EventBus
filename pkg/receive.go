package pkg

import (
	"log"
)

type Dispatcher map[string]func(map[string]any)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
