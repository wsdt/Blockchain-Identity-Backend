package errorHandling

import "log"

func LogErr(err error) {
	if err != nil {
		// Error occurred
		log.Fatal(err)
	}
}
