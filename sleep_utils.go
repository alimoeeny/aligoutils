package aliU

import (
	"math/rand"
	"time"
)

func sleepItOff(d time.Duration) {
	rand.Seed(time.Now().UTC().UnixNano())
	rd := rand.Int63n(int64(d))
	time.Sleep(time.Duration(rd))
}
