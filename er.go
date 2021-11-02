package tim

import (
	"log"
	"time"
)

func RunTime(s string) (string, time.Time) {
	log.Println("Go:	", s)
	return s, time.Now()
}

func TrackThisFunctions(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("Stop:	", s, "took", endTime.Sub(startTime))
}
