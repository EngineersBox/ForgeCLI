package common

import "log"

func CheckError(e error) {
	if e != nil {
		log.Fatalf("[ERROR] %s\n", e)
	}
}
