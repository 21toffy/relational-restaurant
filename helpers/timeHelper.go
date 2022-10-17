package helpers

import (
	"time"
)

func GetCurrentTime() time.Time {
	time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return time
}
