package util

import (
	"log"
	"time"
)

func LogSlowStepTimes(name string, start time.Time) {
	times := time.Now().Sub(start).Seconds()
	if times > 2 {
		log.Printf("%v times:%03vs\n\n", name, times)
	}
}
