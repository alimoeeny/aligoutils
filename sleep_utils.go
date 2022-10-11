package aliU

import (
	"math/rand"
	"time"
)

func SleepItOff(d time.Duration) {
	rand.Seed(time.Now().UTC().UnixNano())
	rd := rand.Int63n(int64(d))
	time.Sleep(time.Duration(rd))
}
