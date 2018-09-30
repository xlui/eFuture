package common

import "log"

// Deal with error.
func FailOnError(e error, message string) {
	if e != nil {
		log.Fatalf("%s: %s", message, e)
	}
}
