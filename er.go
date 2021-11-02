package tim

import (
	"log"
	"time"
)

func RunTime(s string) (string, time.Time) {
	log.Println(s, " start")
	return s, time.Now()
}

func TrackThisFunctions(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println(s, " - ", endTime.Sub(startTime), "\n\n\n\n\n_________________________________")
}
