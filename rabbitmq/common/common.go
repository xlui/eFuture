package common

import "log"

func FailOnError(e error, message string) {
	if e != nil {
		log.Fatalf("%s: %s", message, e)
	}
}
